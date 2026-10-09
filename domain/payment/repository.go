package payment

import (
	"context"

	"goCartAPI/domain/shared"
)

type Repository interface {
	Save(ctx context.Context, transaction *Transaction) error
	FindByID(ctx context.Context, id shared.ID) (*Transaction, error)
	FindByOrderID(ctx context.Context, orderID shared.ID) (*Transaction, error)
	ListByUserID(ctx context.Context, userID shared.ID, pagination shared.Pagination) ([]*Transaction, error)
}
