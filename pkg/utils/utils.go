package utils

import (
	"crypto/rand"
	"math/big"
)

func GeneratePassword(alphabet string, length int) string {
	password := ""
	alphaLen := int64(len(alphabet))
	for len(password) < length {
		n, _ := rand.Int(rand.Reader, big.NewInt(alphaLen-1))
		i := n.Int64()
		password += string(alphabet[i])
	}
	return password
}
