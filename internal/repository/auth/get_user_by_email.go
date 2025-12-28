package authrepo

import (
	"context"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
)

func (r *AuthRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User

	err := r.DB.QueryRow(
		ctx,
		`SELECT id, email, password_hash, created_at, email_verified FROM users WHERE email=$1`,
		email,
	).Scan(
		&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.Verified,
	)

	return user, err
}