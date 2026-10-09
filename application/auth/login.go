package auth

import (
	"context"
	"errors"
	"goCartAPI/domain/user"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

func (s *service) Login(ctx context.Context, in LoginInput) (*LoginResult, error) {
	u, err := s.users.FindByEmail(ctx, in.Email)
	if err != nil {
		if errors.Is(err, user.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if err := s.hasher.Compare(u.PasswordHash(), in.Password); err != nil {
		return nil, ErrInvalidCredentials
	}

	access, err := s.tokens.GenerateAccessToken(u.ID(), u.Role())
	if err != nil {
		return nil, err
	}

	refresh, err := s.tokens.GenerateRefreshToken(u.ID(), u.Role())
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		User:         u,
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil

}
