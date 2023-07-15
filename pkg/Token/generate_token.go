package token

import (
	"crypto/rand"
	"encoding/hex"
)

func GeneratedToken() (string, error) {
	byteToken := make([]byte, 16)
	_, err := rand.Read(byteToken)

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(byteToken), nil
}
