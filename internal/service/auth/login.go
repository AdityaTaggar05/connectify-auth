package authservice

import (
	"context"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, email, password string) (model.TokenPair, error) {
	if !isValidEmail(email) {
		return model.TokenPair{}, ErrInvalidEmailFormat
	}

	if password == "" {
		return model.TokenPair{}, ErrInvalidPasswordFormat
	}

	tokens := model.TokenPair{}

	user, err := s.AuthRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return tokens, ErrUserNotFound
	}

	if !user.Verified {
		return tokens, ErrEmailNotVerified
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return tokens, ErrIncorrectPassword
	}

	tokens.AccessToken, err = model.GenerateJWT(user.ID, s.SigningKey, s.Config.AccessTTL)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := model.GenerateRefreshToken(user.ID, s.Config.RefreshTTL)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = refreshToken.Token

	err = s.TokenRepo.CreateRefreshToken(ctx, user.ID, tokens.RefreshToken, time.Now().Add(s.Config.RefreshTTL))
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}
