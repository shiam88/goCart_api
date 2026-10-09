package admin

import (
	"context"
	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
)

func (s *service) UpdateOrderStatus(ctx context.Context, id shared.ID, status order.Status) (*order.Order, error) {
	return s.orderSvc.UpdateStatus(ctx, id, status)
}
