package cart

import (
	"context"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
)

func (s *service) Clear(ctx context.Context, userID shared.ID) (*cart.Cart, error) {
	c, err := s.carts.FindByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	c.Clear()
	if err := s.carts.Save(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
