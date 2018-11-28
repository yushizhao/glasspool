package eth

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
)

func Hash(b []byte) string {
	sha256sum := sha256.Sum256(b)
	return hex.EncodeToString(sha256sum[:])
}

func HashInt(i int64) string {
	b := itob(i)
	return Hash(b)
}

// modified from glasspool common util
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
