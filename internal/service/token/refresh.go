package tokenservice

import (
	"context"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
)

func (s *Service) Refresh(ctx context.Context, oldToken string) (model.TokenPair, error) {
	if !IsValidRefreshToken(oldToken) {
		return model.TokenPair{}, ErrInvalidRefreshTokenFormat
	}

	tokens := model.TokenPair{}

	rt, err := s.TokenRepo.GetRefreshToken(ctx, oldToken)

	if err != nil || rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		return tokens, ErrInvalidRefreshToken
	}

	user, err := s.AuthRepo.GetUserByID(ctx, rt.UserID)

	err = s.TokenRepo.RevokeRefreshToken(ctx, oldToken)
	if err != nil {
		return tokens, err
	}

	refreshToken, err := model.GenerateRefreshToken(rt.UserID, s.Config.RefreshTTL)
	if err != nil {
		return tokens, err
	}
	tokens.RefreshToken = refreshToken.Token

	_ = s.TokenRepo.CreateRefreshToken(ctx, rt.UserID, tokens.RefreshToken, time.Now().Add(s.Config.RefreshTTL))
	tokens.AccessToken, err = model.GenerateJWT(user, s.SigningKey, s.Config.AccessTTL)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}