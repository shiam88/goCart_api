package order

import (
	"context"
	"errors"
	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
)

var ErrOrderForbidden = errors.New("order does not belong to authenticated user")

func (s *service) Get(ctx context.Context, id, userID shared.ID, admin bool) (*order.Order, error) {
	o, err := s.orders.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if !admin && o.UserID() != userID {
		return nil, ErrOrderForbidden
	}
	return o, nil
}
