package tokenrepo

import (
	"context"
	"time"
)

func (r *TokenRepository) CreateRefreshToken(ctx context.Context, user_id, token string, exp time.Time) error {
	_, err := r.DB.Exec(
		ctx,
		`INSERT INTO refresh_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)`,
		user_id, token, exp,
	)
	return err
}
