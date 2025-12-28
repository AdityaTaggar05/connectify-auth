package authrepo

import "context"

func (r *AuthRepository) CreateUser(ctx context.Context, email, hash string) error {
	_, err := r.DB.Exec(
		ctx,
		`INSERT INTO users (email, password_hash) VALUES ($1, $2)`,
		email, hash,
	)

	return err
}
