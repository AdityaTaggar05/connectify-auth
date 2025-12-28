package tokenrepo

import (
	"context"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
)

func (r *TokenRepository) GetRefreshToken(ctx context.Context, token string) (model.RefreshToken, error) {
	var rt model.RefreshToken

	err := r.DB.QueryRow(
		ctx,
		`SELECT user_id, token, revoked, expires_at FROM refresh_tokens WHERE token=$1`,
		token,
	).Scan(
		&rt.UserID, &rt.Token, &rt.Revoked, &rt.ExpiresAt,
	)
	return rt, err
}
