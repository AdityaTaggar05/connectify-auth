package tokenrepo

import "context"

func (r *TokenRepository) RevokeRefreshToken(ctx context.Context, token string) error {
	_, err := r.DB.Exec(
		ctx,
		`UPDATE refresh_tokens SET revoked=true WHERE token=$1`,
		token,
	)
	return err
}
