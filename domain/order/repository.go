package order

import (
	"context"
	"goCartAPI/domain/shared"
)

type Repository interface {
	Save(ctx context.Context, order *Order) error
	FindByID(ctx context.Context, id shared.ID) (*Order, error)
	ListByUserID(ctx context.Context, userID shared.ID, pagination shared.Pagination) ([]*Order, error)
	List(ctx context.Context, pagination shared.Pagination) ([]*Order, error)
	CountByUserID(ctx context.Context, userID shared.ID) (int, error)
}
