package inventory

import (
	"context"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/shared"
)

type Service interface {
	Get(ctx context.Context, productID shared.ID) (*inventory.Inventory, error)
	Adjust(ctx context.Context, productID shared.ID, delta int) (*inventory.Inventory, error)
}
