package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/models"
)

type Config struct {
	PORT string
	DB_URL     string
	JWT_SIGNING_KEY *models.SigningKey
	JWT_EXP    time.Duration
	REFRESH_EXP time.Duration
}

func Load() Config {
	jwt_exp, _ := strconv.Atoi(os.Getenv("JWT_EXP"))
	refresh_exp, _ := strconv.Atoi(os.Getenv("REFRESH_EXP"))
	key, err := loadSigningKey()

	if err != nil {
		panic(errors.New("failed to load JWT private key: " + err.Error()))
	}

	return Config{
		PORT: os.Getenv("PORT"),
		DB_URL: os.Getenv("DATABASE_URL"),
		JWT_SIGNING_KEY: key,
		JWT_EXP: time.Duration(jwt_exp) * time.Minute,
		REFRESH_EXP: time.Duration(refresh_exp) * (time.Hour * 24),
	}
}

func loadSigningKey() (*models.SigningKey, error) {
	keyData, err := os.ReadFile(os.Getenv("JWT_PRIVATE_KEY"))

	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode([]byte(keyData))
	if block == nil {
		return nil, errors.New("invalid private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privateKey := key.(*rsa.PrivateKey)

	return &models.SigningKey{
		ID: "auth-key-2025-01",
		PrivateKey: privateKey,
		PublicKey: &privateKey.PublicKey,
	}, nil
}