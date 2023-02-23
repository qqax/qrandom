package qrandom

import (
	"crypto/rand"
	"errors"
	"net/url"
)

// CreateEscapedString create a random escaped string of length n, so it can be safely placed inside a URL query.
func CreateEscapedString(n int) (string, error) {
	randArray := make([]byte, n)
	if _, err := rand.Read(randArray); err != nil {
		return "", errors.New("unable to create random url")
	}
	return url.QueryEscape(string(randArray)), nil
}

//func RandomString(n int) (string, error) {
//	randArray := make([]byte, n)
//	if _, err := rand.Read(randArray); err != nil {
//		return "", errors.New("unable to create random string")
//	}
//	return string(randArray), nil
//}
