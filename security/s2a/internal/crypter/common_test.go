package crypter

import (
	"bytes"
	"google.golang.org/grpc/security/s2a/internal/crypter/testutil"
	"testing"
)

// fakeAEAD is a fake implementation of an AEAD interface used for testing.
type fakeAEAD struct{}

func (*fakeAEAD) NonceSize() int                                  { return nonceSize }
func (*fakeAEAD) Overhead() int                                   { return tagSize }
func (*fakeAEAD) Seal(_, _, plaintext, _ []byte) []byte           { return plaintext }
func (*fakeAEAD) Open(_, _, ciphertext, _ []byte) ([]byte, error) { return ciphertext, nil }

func TestSliceForAppend(t *testing.T) {
	for _, tc := range []struct {
		desc  string
		inBuf []byte
		n     int
	}{
		{
			desc: "nil buf and zero length",
		},
		{
			desc: "nil buf and non-zero length",
			n:    5,
		},
		{
			desc:  "non-empty buf and zero length",
			inBuf: testutil.Dehex("1111111111"),
		},
		{
			desc:  "non-empty buf and non-zero length",
			inBuf: testutil.Dehex("1111111111"),
			n:     5,
		},
		{
			desc:  "test slice capacity pre allocated",
			inBuf: make([]byte, 0, 5),
			n:     5,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			head, tail := sliceForAppend(tc.inBuf, tc.n)
			// Check that the resulting head buffer starts with the same byte
			// sequence as the input buffer.
			if got, want := head, tc.inBuf; !bytes.HasPrefix(head, tc.inBuf) {
				t.Errorf("sliceForAppend(%v, %v).head = %v, want %v", tc.inBuf, tc.n, got, want)
			}
			// Check that the length of the resulting head buffer is equal
			// to the initial buffer + the additional length requested.
			if got, want := len(head), len(tc.inBuf)+tc.n; got != want {
				t.Errorf("sliceForAppend(%v, %v).tail = %v, want %v", tc.inBuf, tc.n, got, want)
			}
			// Check that the length of the resulting tail buffer is what was
			// requested.
			if got, want := len(tail), tc.n; got != want {
				t.Errorf("sliceForAppend(%v, %v).tail = %v, want %v", tc.inBuf, tc.n, got, want)
			}
		})
	}
}

func TestInvalidNonceSize(t *testing.T) {
	nonce := []byte("1")
	if _, err := encrypt(&fakeAEAD{}, nil, nil, nonce, nil); err == nil {
		t.Errorf("encrypt(&fakeAEAD{}, nil, nil, %v, nil) expected error, received none", nonce)
	}
	if _, err := decrypt(&fakeAEAD{}, nil, nil, nonce, nil); err == nil {
		t.Errorf("decrypt(&fakeAEAD{}, nil, nil, %v, nil) expected error, received none", nonce)
	}
}

func TestEncrypt(t *testing.T) {
	plaintext := []byte("test")
	nonce := make([]byte, nonceSize)
	ciphertext, err := decrypt(&fakeAEAD{}, nil, plaintext, nonce, nil)
	if err != nil {
		t.Fatalf("encrypt(&fakeAEAD{}, nil, %v, %v, nil) failed: %v", plaintext, nonce, err)
	}
	if got, want := ciphertext, plaintext; !bytes.Equal(got, want) {
		t.Fatalf("encrypt(&fakeAEAD{}, nil, %v, %v, nil) = %v, want %v", plaintext, nonce, got, want)
	}
}

func TestDecrypt(t *testing.T) {
	ciphertext := []byte("test")
	nonce := make([]byte, nonceSize)
	plaintext, err := decrypt(&fakeAEAD{}, nil, ciphertext, nonce, nil)
	if err != nil {
		t.Fatalf("decrypt(&fakeAEAD{}, nil, %v, %v, nil) failed: %v", ciphertext, nonce, err)
	}
	if got, want := plaintext, ciphertext; !bytes.Equal(got, want) {
		t.Fatalf("decrypt(&fakeAEAD{}, nil, %v, %v, nil) = %v, want %v", ciphertext, nonce, got, want)
	}
}
