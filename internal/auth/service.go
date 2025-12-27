package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/config"
	"github.com/AdityaTaggar05/annora-auth/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepo      *UserRepository
	RefreshRepo *RefreshTokenRepository
	Config config.Config
}


func NewService(DB *pgxpool.Pool, cfg config.Config) *Service {
	return &Service{
		UserRepo: &UserRepository{DB: DB},
		RefreshRepo: &RefreshTokenRepository{DB: DB},
		Config: cfg,
	}
}

func (s *Service) Register(ctx context.Context, email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	return s.UserRepo.CreateUser(ctx, email, string(hash))
}

func (s *Service) Login(ctx context.Context, email, password string) (Tokens, error) {
	tokens := Tokens{}

	user, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return tokens, err
	}
	
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return tokens, err
	}

	tokens.AccessToken, err = utils.GenerateJWT(user.ID, s.Config.JWT_SIGNING_KEY, s.Config.JWT_EXP)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = utils.GenerateRefreshToken()
	if err != nil {
		return tokens, err
	}
	
	err = s.RefreshRepo.Create(ctx, user.ID, tokens.RefreshToken, time.Now().Add(s.Config.REFRESH_EXP))
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func (s *Service) Refresh(ctx context.Context, oldToken string) (Tokens, error) {
	tokens := Tokens{}
	
	rt, err := s.RefreshRepo.Get(ctx, oldToken)
	if err != nil || rt.Revoked || rt.ExpiresAt.Before(time.Now()) {
		return tokens, fmt.Errorf("Unauthorized Request")
	}

	err = s.RefreshRepo.Revoke(ctx, oldToken)
	if err != nil {
		return tokens, err
	}

	tokens.RefreshToken, err = utils.GenerateRefreshToken()
	if err != nil {
		return tokens, err
	}
	_ = s.RefreshRepo.Create(ctx, rt.UserID, tokens.RefreshToken, time.Now().Add(s.Config.REFRESH_EXP))

	tokens.AccessToken, err = utils.GenerateJWT(rt.UserID, s.Config.JWT_SIGNING_KEY, s.Config.JWT_EXP)
	if err != nil {
		return tokens, err
	}

	return tokens, nil
}

func (s *Service) Logout(ctx context.Context, oldToken string) error {
	return s.RefreshRepo.Revoke(ctx, oldToken)
}