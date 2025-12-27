package utils

import (
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id string, signingKey *models.SigningKey, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.RegisteredClaims{
		Subject: id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Issuer: "annora-auth",
	})

	token.Header["kid"] = signingKey.ID

	return token.SignedString(signingKey.PrivateKey)
}