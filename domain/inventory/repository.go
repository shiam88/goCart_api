package inventory

import (
	"context"

	"goCartAPI/domain/shared"
)

type Repository interface {
	Save(ctx context.Context, inventory *Inventory) error
	FindByProductID(ctx context.Context, productID shared.ID) (*Inventory, error)
}
