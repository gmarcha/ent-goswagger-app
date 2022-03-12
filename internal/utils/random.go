package utils

import (
	"math/rand"
	"time"
)

const charList = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomString returns a random generated string.
func RandomString(size int) string {
	b := make([]byte, size)
	for i := range b {
		b[i] = charList[rand.Intn(len(charList))]
	}
	return string(b)
}
