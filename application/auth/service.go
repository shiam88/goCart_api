package auth

import "goCartAPI/domain/user"

type service struct {
	users  user.Repository
	hasher PasswordHasher
	tokens TokenService
}

func NewService(users user.Repository, hasher PasswordHasher, tokens TokenService) Service {
	return &service{
		users:  users,
		hasher: hasher,
		tokens: tokens,
	}
}
