package cart

import (
	"context"
	"goCartAPI/domain/shared"
)

type Repository interface {
	Save(ctx context.Context, cart *Cart) error
	FindByID(ctx context.Context, id shared.ID) (*Cart, error)
	FindByUserID(ctx context.Context, userID shared.ID) (*Cart, error)
	Delete(ctx context.Context, id shared.ID) error
}
