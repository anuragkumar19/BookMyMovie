package storage

import (
	"crypto/rand"
	"encoding/base64"
)

func randObjKey() (string, error) {
	b := make([]byte, 64)

	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}
