package authrepo

import (
	"context"

	"github.com/AdityaTaggar05/annora-auth/internal/model"
)

func (r *AuthRepository) GetUserByID(ctx context.Context, id string) (model.User, error) {
	var user model.User

	err := r.DB.QueryRow(
		ctx,
		`SELECT id, email, role, password_hash, created_at, email_verified FROM users WHERE id=$1`,
		id,
	).Scan(
		&user.ID, &user.Email, &user.Role, &user.Password, &user.CreatedAt, &user.Verified,
	)

	return user, err
}