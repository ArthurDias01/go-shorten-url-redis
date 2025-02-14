package api

import "math/rand/v2"

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(db map[string]string) string {
	const n = 8
	for {
		bytes := make([]byte, n)
		for i := range n {
			bytes[i] = characters[rand.IntN(len(characters))]
		}
		code := string(bytes)
		if _, exists := db[code]; !exists {
			return code
		}
	}
}
