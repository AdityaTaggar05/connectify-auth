package tokenservice

import (
	"github.com/AdityaTaggar05/annora-auth/internal/config"
	"github.com/AdityaTaggar05/annora-auth/internal/model"
	authrepo "github.com/AdityaTaggar05/annora-auth/internal/repository/auth"
	tokenrepo "github.com/AdityaTaggar05/annora-auth/internal/repository/token"
)

type Service struct {
	AuthRepo   *authrepo.AuthRepository
	TokenRepo  *tokenrepo.TokenRepository
	Config     config.JWTConfig
	SigningKey *model.SigningKey
}

func NewService(authRepo *authrepo.AuthRepository, tokenRepo *tokenrepo.TokenRepository, cfg config.JWTConfig, signingKey *model.SigningKey) *Service {
	return &Service{
		AuthRepo: authRepo,
		TokenRepo: tokenRepo,
		Config: cfg,
		SigningKey: signingKey,
	}
}
