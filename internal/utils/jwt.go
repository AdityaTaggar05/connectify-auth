package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id, secret string, exp int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":id,
		"exp":time.Now().Add(time.Hour * 24 * time.Duration(exp)).Unix(),
	})

	return token.SignedString(secret)
}