package cart

import (
	"context"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
)

func (s *service) Get(ctx context.Context, userID shared.ID) (*cart.Cart, error) {
	return s.carts.FindByUserID(ctx, userID)
}
