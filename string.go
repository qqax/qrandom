package qrandom

import (
	"math/rand"
	"time"
	"unsafe"
)

// This code was found on https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go,
// all rights belong to https://stackoverflow.com/users/1705598/icza

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// String returns the accumulated string.
//func (b *Builder) String() string {
//	return *(*string)(unsafe.Pointer(&b.buf))
//}

// String Random string from bytes with mask, improved with source and unsafe.
func String(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}
