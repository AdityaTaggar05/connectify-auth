package utils

import "crypto/rand"

func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)

	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	
	return string(b), nil
}