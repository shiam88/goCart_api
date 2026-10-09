package auth

import (
	"context"
	"goCartAPI/domain/shared"
	"goCartAPI/domain/user"
)

type Service interface {
	Register(ctx context.Context, in RegisterInput) (*user.User, error)
	Login(ctx context.Context, in LoginInput) (*LoginResult, error)
	Refresh(ctx context.Context, refreshToken string) (*RefreshResult, error)
	Logout(ctx context.Context) error
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}

type TokenService interface {
	GenerateAccessToken(userID shared.ID, role user.Role) (string, error)
	GenerateRefreshToken(userID shared.ID, role user.Role) (string, error)
	ParseAccessToken(token string) (TokenClaims, error)
	ParseRefreshToken(token string) (TokenClaims, error)
}

type TokenClaims struct {
	UserID shared.ID
	Role   user.Role
}

type RegisterInput struct {
	Name     string
	Email    string
	Password string
}

type LoginInput struct {
	Email    string
	Password string
}

type LoginResult struct {
	User         *user.User
	AccessToken  string
	RefreshToken string
}

type RefreshResult struct {
	AccessToken string
}
