package auth

import (
	"context"
	"time"

	"github.com/AdityaTaggar05/annora-auth/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RefreshTokenRepository struct {
	DB *pgxpool.Pool
}

func (r *RefreshTokenRepository) Create(ctx context.Context, user_id, token string, exp time.Time) error {
	_, err := r.DB.Exec(ctx,
	`INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`,
	user_id, token, exp)
	return err
}

func (r *RefreshTokenRepository) Get(ctx context.Context, token string) (models.RefreshToken, error) {
	var rt models.RefreshToken

	err := r.DB.QueryRow(ctx,
	`SELECT user_id, token, revoked, expires_at FROM refresh_tokens WHERE token=$1`,
	token).Scan(&rt.UserID, &rt.Token, &rt.Revoked, &rt.ExpiresAt)
	return rt, err
}

func (r *RefreshTokenRepository) Revoke(ctx context.Context, token string) error {
	_, err := r.DB.Exec(ctx,
	`UPDATE refresh_tokens SET revoked=true WHERE token=$1`,
	token)
	return err
}
