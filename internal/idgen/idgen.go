package idgen

import (
	"crypto/rand"
	"math/big"
)

// contains only lower case letters and numbers
func GenerateRandomId(length int) (string, error) {
	const set = "abcdefghijklmnopqrstuvwxyz0123456789"
	ret := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(set))))
		if err != nil {
			return "", err
		}
		ret[i] = set[num.Int64()]
	}
	return string(ret), nil
}
