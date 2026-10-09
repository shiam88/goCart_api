package order

import (
	"context"

	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
)

func (s *service) UpdateStatus(
	ctx context.Context,
	orderID shared.ID,
	target order.Status,
) (*order.Order, error) {

	o, err := s.orders.FindByID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	if err := o.TransitionTo(target); err != nil {
		return nil, err
	}

	if err := s.orders.Save(ctx, o); err != nil {
		return nil, err
	}

	return o, nil
}
