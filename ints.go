package qrandom

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func RandIntBetween(min, max int64) (int64, error) {
	if min >= max {
		return 0, fmt.Errorf("max should be more as min, got min = %d & max = %d", min, max)
	}

	randInt, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		panic(err)
	}

	return min + randInt.Int64(), nil
}
