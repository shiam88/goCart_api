package admin

import (
	"context"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/shared"
)

func (s *service) AdjustStock(ctx context.Context, productID shared.ID, delta int) (*inventory.Inventory, error) {
	return s.stockSvc.Adjust(ctx, productID, delta)
}
