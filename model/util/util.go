package util

import (
	"encoding/base64"
	"golang.org/x/crypto/sha3"
)

func Sha(input string) string {
	if input == "" {
		return ""
	}
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(input))
	return base64.StdEncoding.EncodeToString(h)
}
