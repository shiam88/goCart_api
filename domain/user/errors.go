package user

import "errors"

var (
	ErrInvalidUserID       = errors.New("invalid user id")
	ErrInvalidName         = errors.New("user name is required")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrMissingPasswordHash = errors.New("password hash is required")
	ErrInvalidRole         = errors.New("invalid user role")
	ErrEmailAlreadyUsed    = errors.New("email is already in use")
	ErrUserNotFound        = errors.New("user not found ")
)
