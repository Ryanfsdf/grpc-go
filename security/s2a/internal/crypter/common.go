package crypter

const (
	// GcmTagSize is the tag size in bytes. From crypto/cipher/gcm.go in the Go
	// crypto library.
	GcmTagSize = 16
	// NonceSize is the size of the nonce in number of bytes for
	// AES-128-GCM-SHA256, AES-256-GCM-SHA384, and CHACHA20-POLY1305-SHA256.
	NonceSize = 12
)

// SliceForAppend takes a slice and a requested number of bytes. It returns a
// slice with the contents of the given slice followed by that many bytes and a
// second slice that aliases into it and contains only the extra bytes. If the
// original slice has sufficient capacity then no allocation is performed.
func SliceForAppend(in []byte, n int) (head, tail []byte) {
	if total := len(in) + n; cap(in) >= total {
		head = in[:total]
	} else {
		head = make([]byte, total)
		copy(head, in)
	}
	tail = head[len(in):]
	return head, tail
}