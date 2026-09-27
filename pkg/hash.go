package kit

import (
	"crypto/sha256"
	"encoding/hex"
)

// SHA256Hex returns the hex-encoded SHA-256 digest of s.
func SHA256Hex[T string | []byte](s T) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}
