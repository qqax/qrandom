package qrandom

import (
	"crypto/rand"
	"errors"
)

func ByteArray(n int) ([]byte, error) {
	randArray := make([]byte, n)
	if _, err := rand.Read(randArray); err != nil {
		return nil, errors.New("unable to create random byte array")
	}
	return randArray, nil
}
