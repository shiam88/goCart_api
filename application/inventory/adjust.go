package inventory

import (
	"context"
	"errors"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/shared"
)

var ErrNegativeAdjustmentUnsupported = errors.New("negative manual adjustment is not supported; use domain stock flows")

func (s *service) Adjust(ctx context.Context, productID shared.ID, delta int) (*inventory.Inventory, error) {
	if delta <= 0 {
		return nil, ErrNegativeAdjustmentUnsupported
	}
	i, err := s.repo.FindByProductID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if err := i.AddStock(delta); err != nil {
		return nil, err
	}
	if err := s.repo.Save(ctx, i); err != nil {
		return nil, err
	}
	return i, nil
}
