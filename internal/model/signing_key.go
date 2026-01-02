package model

import (
	"crypto/rsa"
	"encoding/base64"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SigningKey struct {
	ID         string
	Issuer     string
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (s *SigningKey) PublicKeyToJWK() map[string]string {
	n := base64.RawURLEncoding.EncodeToString(s.PublicKey.N.Bytes())
    e := base64.RawURLEncoding.EncodeToString(big.NewInt(int64(s.PublicKey.E)).Bytes())

    return map[string]string{
        "kty": "RSA",
        "kid": s.ID,
        "use": "sig",
        "alg": "RS256",
        "n":   n,
        "e":   e,
    }
}

func GenerateJWT(user User, signingKey *SigningKey, ttl time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"sub": user.ID,
		"role": user.Role,
		"exp": jwt.NewNumericDate(time.Now().Add(ttl)),
		"iat": jwt.NewNumericDate(time.Now()),
		"iss": signingKey.Issuer,
	})

	token.Header["kid"] = signingKey.ID

	return token.SignedString(signingKey.PrivateKey)
}
