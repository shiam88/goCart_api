package inventory

import (
	"context"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/shared"
)

func (s *service) Get(ctx context.Context, productID shared.ID) (*inventory.Inventory, error) {
	return s.repo.FindByProductID(ctx, productID)
}
