package qrandom

import (
	"crypto/rand"
	"errors"
)

// CreateByteArray creates a random byte array of length n.
func CreateByteArray(n int) ([]byte, error) {
	randArray := make([]byte, n)
	if _, err := rand.Read(randArray); err != nil {
		return nil, errors.New("unable to create random byte array")
	}
	return randArray, nil
}
