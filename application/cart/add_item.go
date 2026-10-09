package cart

import (
	"context"
	"errors"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
)

func (s *service) AddItem(ctx context.Context, userID, productID shared.ID, qty int) (*cart.Cart, error) {
	p, err := s.products.FindByID(ctx, productID)
	if err != nil {
		return nil, err
	}
	if !p.IsActive() {
		return nil, errors.New("product is inactive")
	}
	c, err := s.carts.FindByUserID(ctx, userID)
	if err != nil {
		if errors.Is(err, cart.ErrCartNotFound) {
			c, err = cart.NewCart(userID, s.currency)
		}
		if err != nil {
			return nil, err
		}
	}
	if err := c.AddItem(p.ID(), p.Name(), p.Price(), qty); err != nil {
		return nil, err
	}
	if err := s.carts.Save(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
