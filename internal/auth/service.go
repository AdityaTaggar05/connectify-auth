package auth

import (
	"context"

	"github.com/AdityaTaggar05/connectify-auth/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo      *Repository
	JWTSecret string
	JWTExp int
}

func (s *Service) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	return s.Repo.CreateUser(ctx, email, string(hash))
}

func (s *Service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.Repo.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.ID, s.JWTSecret, s.JWTExp)
}