package auth

import (
	"context"
	"errors"
	"strings"

	"goCartAPI/domain/user"
)

func (s *service) Register(ctx context.Context, in RegisterInput) (*user.User, error) {
	if len(strings.TrimSpace(in.Password)) < 8 {
		return nil, errors.New("password must be at least 8 characters")
	}
	email := strings.ToLower(strings.TrimSpace(in.Email))
	exists, err := s.users.EmailExists(ctx, email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, user.ErrEmailAlreadyUsed
	}

	hash, err := s.hasher.Hash(in.Password)
	if err != nil {
		return nil, err
	}

	u, err := user.NewUser(in.Name, email, hash)
	if err != nil {
		return nil, err
	}
	if err := s.users.Save(ctx, u); err != nil {
		return nil, err
	}
	return u, nil
}
