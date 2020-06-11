package crypter

import (
	"fmt"
	"golang.org/x/crypto/cryptobyte"
	s2a_proto "google.golang.org/grpc/security/s2a/internal/proto"
	"hash"
	"sync"
)

// uint64Size is the size of a uint64 in bytes.
const uint64Size = 8

const (
	tls13Key    = "tls13 key"
	tls13Nonce  = "tls13 iv"
	tls13Update = "tls13 traffic upd"
)

type S2AHalfConnection struct {
	cs            ciphersuite
	h             func() hash.Hash
	aeadCrypter   s2aAeadCrypter
	expander      hkdfExpander
	seqCounter    counter
	mutex         sync.Mutex
	trafficSecret []byte
	nonce         []byte
	key           []byte
}

// NewHalfConn creates a new instance of S2AHalfConnection.
func NewHalfConn(ciphersuite s2a_proto.Ciphersuite, trafficSecret []byte) (S2AHalfConnection, error) {
	cs := newCiphersuite(ciphersuite)
	if cs.trafficSecretSize() != len(trafficSecret) {
		return S2AHalfConnection{}, fmt.Errorf("supplied traffic secret must be %v bytes, given: %v", cs.trafficSecretSize(), trafficSecret)
	}

	hc := S2AHalfConnection{cs: cs, h: cs.hashFunction(), expander: &defaultHKDFExpander{}, seqCounter: newCounter(), trafficSecret: trafficSecret}

	var err error
	if err = hc.updateWithNewTrafficSecret(hc.trafficSecret); err != nil {
		return S2AHalfConnection{}, fmt.Errorf("hc.updateWithNewTrafficSecret(%v) failed with error: %v", hc.trafficSecret, err)
	}

	hc.aeadCrypter, err = cs.aeadCrypter(hc.key)
	if err != nil {
		return S2AHalfConnection{}, fmt.Errorf("cs.aeadCrypter(%v) failed with error: %v", hc.key, err)
	}
	return hc, nil
}

// Encrypt encrypts the plaintext and computes the tag of dst and plaintext.
// dst and plaintext may fully overlap or not at all. Note that the sequence
// number will still be incremented on failure, unless the sequence has
// overflowed.
func (hc *S2AHalfConnection) Encrypt(dst, plaintext, aad []byte) ([]byte, error) {
	hc.mutex.Lock()
	sequence, err := hc.getAndIncrementSequence()
	if err != nil {
		hc.mutex.Unlock()
		return nil, err
	}
	nonce := hc.maskedNonce(sequence)
	crypter := hc.aeadCrypter
	hc.mutex.Unlock()
	return crypter.encrypt(dst, plaintext, nonce, aad)
}

// Decrypt decrypts ciphertext and verifies the tag. dst and ciphertext may
// fully overlap or not at all. Note that the sequence number will still be
// incremented on failure, unless the sequence has overflowed.
func (hc *S2AHalfConnection) Decrypt(dst, ciphertext, aad []byte) ([]byte, error) {
	hc.mutex.Lock()
	sequence, err := hc.getAndIncrementSequence()
	if err != nil {
		hc.mutex.Unlock()
		return nil, err
	}
	nonce := hc.maskedNonce(sequence)
	crypter := hc.aeadCrypter
	hc.mutex.Unlock()
	return crypter.decrypt(dst, ciphertext, nonce, aad)
}

// UpdateKey updates the traffic secret key, as specified in
// https://tools.ietf.org/html/rfc8446#section-7.2
func (hc *S2AHalfConnection) UpdateKey() error {
	hc.mutex.Lock()
	defer hc.mutex.Unlock()

	var err error
	hc.trafficSecret, err = hc.deriveSecret(hc.trafficSecret, []byte(tls13Update), hc.cs.trafficSecretSize())
	if err != nil {
		return fmt.Errorf("hc.deriveSecret(h, %v, %v, %v) failed with error: %v", hc.trafficSecret, tls13Update, hc.cs.trafficSecretSize(), err)
	}

	if err = hc.updateWithNewTrafficSecret(hc.trafficSecret); err != nil {
		return fmt.Errorf("hc.updateWithNewTrafficSecret(%v) failed with error: %v", hc.trafficSecret, err)
	}

	err = hc.aeadCrypter.updateKey(hc.key)
	if err != nil {
		return fmt.Errorf("hc.aeadCrypter.updateKey(%v) failed with error: %v", hc.key, err)
	}

	hc.seqCounter.reset()
	return nil
}

// updateWithNewTrafficSecret takes a new traffic secret and updates the key
// and nonce.
func (hc *S2AHalfConnection) updateWithNewTrafficSecret(newTrafficSecret []byte) error {
	var err error
	hc.key, err = hc.deriveSecret(newTrafficSecret, []byte(tls13Key), hc.cs.keySize())
	if err != nil {
		return fmt.Errorf("hc.deriveSecret(h, %v, %v, %v) failed with error: %v", hc.trafficSecret, tls13Key, hc.cs.keySize(), err)
	}

	hc.nonce, err = hc.deriveSecret(newTrafficSecret, []byte(tls13Nonce), hc.cs.nonceSize())
	if err != nil {
		return fmt.Errorf("hc.deriveSecret(h, %v, %v, %v) failed with error: %v", hc.trafficSecret, tls13Nonce, hc.cs.nonceSize(), err)
	}
	return nil
}

// getAndIncrement returns the current sequence number and increments it.
func (hc *S2AHalfConnection) getAndIncrementSequence() (uint64, error) {
	sequence, err := hc.seqCounter.value()
	if err != nil {
		return 0, err
	}
	hc.seqCounter.increment()
	return sequence, nil
}

// maskedNonce creates a new S2A nonce using the sequence number.
func (hc *S2AHalfConnection) maskedNonce(sequence uint64) []byte {
	nonce := make([]byte, len(hc.nonce))
	copy(nonce, hc.nonce)
	for i := 0; i < uint64Size; i++ {
		nonce[nonceSize-uint64Size+i] ^= byte(sequence >> uint64(56-uint64Size*i))
	}
	return nonce
}

// deriveSecret implements Derive-Secret specified in
// https://tools.ietf.org/html/rfc8446#section-7.1.
func (hc *S2AHalfConnection) deriveSecret(secret, label []byte, length int) ([]byte, error) {
	var hkdfLabel cryptobyte.Builder
	hkdfLabel.AddUint16(uint16(length))
	hkdfLabel.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
		b.AddBytes(label)
	})
	// Append empty `Context` field, specified in the RFC. The Half Connection
	// does not use the `Context` field.
	hkdfLabel.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
		b.AddBytes([]byte(""))
	})
	hkdfLabelBytes, err := hkdfLabel.Bytes()
	if err != nil {
		return nil, fmt.Errorf("deriveSecret failed with error: %v", err)
	}
	return hc.expander.expand(hc.h, secret, hkdfLabelBytes, length)
}
