package shortener

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortCode() (string, error) {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}

	return base64.URLEncoding.EncodeToString(b), nil
}
