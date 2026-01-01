package authservice

import "errors"

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrInvalidPasswordFormat = errors.New("invalid password format")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrInvalidToken = errors.New("token has expired or has been already used")
	ErrUserNotFound = errors.New("user not found")
	ErrEmailNotVerified = errors.New("email not verified")
	ErrIncorrectPassword = errors.New("incorrect password")
)
