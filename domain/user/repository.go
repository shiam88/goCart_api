package user

import (
	"context"

	"goCartAPI/domain/shared"
)

// Repository is a domain contract. Its PostgreSQL implementation belongs to infrastructure.
type Repository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id shared.ID) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	EmailExists(ctx context.Context, email string) (bool, error)
}
