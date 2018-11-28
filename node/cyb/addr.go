package cyb

import (
	"encoding/base64"
	"time"
)

// concat a random string
func GenerateAddress() string {
	t := time.Now().String()
	s := base64.StdEncoding.EncodeToString([]byte(t))
	return ACCOUNT + s
}
