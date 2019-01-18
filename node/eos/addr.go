package eos

import (
	"crypto/sha256"
	"encoding/base64"
	"time"
)

// concat a random string
func GenerateAddress() string {
	t := time.Now().String()
	sha256sum := sha256.Sum256([]byte(t))
	s := base64.StdEncoding.EncodeToString(sha256sum[20:])
	return ACCOUNT + "[" + s + "]"
}
