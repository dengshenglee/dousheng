package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
