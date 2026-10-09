package cart

import (
	"context"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
)

func (s *service) RemoveItem(ctx context.Context, userID, productID shared.ID) (*cart.Cart, error) {
	c, err := s.carts.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if err := c.RemoveItem(productID); err != nil {
		return nil, err
	}
	if err := s.carts.Save(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
