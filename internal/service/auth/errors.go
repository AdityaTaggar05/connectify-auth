package authservice

import "errors"

var (
	ErrEmailNotVerified = errors.New("email not verified")
)