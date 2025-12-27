package models

import "crypto/rsa"

type SigningKey struct {
	ID         string
	PrivateKey *rsa.PrivateKey
	PublicKey *rsa.PublicKey
}
