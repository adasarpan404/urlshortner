package utils

import (
	"math/rand"
	"time"
)

func GenerateShortCode() string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}
