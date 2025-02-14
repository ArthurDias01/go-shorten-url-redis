package store

import (
	"math/rand/v2"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode() string {
	const n = 8
	bytes := make([]byte, n)
	for i := range n {
		bytes[i] = characters[rand.IntN(len(characters))]
	}
	code := string(bytes)
	return code
}
