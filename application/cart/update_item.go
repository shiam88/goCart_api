package cart

import (
	"context"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
)

func (s *service) UpdateItem(ctx context.Context, userID, productID shared.ID, qty int) (*cart.Cart, error) {
	c, err := s.carts.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if err := c.UpdateQuantity(productID, qty); err != nil {
		return nil, err
	}
	if err := s.carts.Save(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
