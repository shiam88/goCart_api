package admin

import (
	"context"
	"goCartAPI/domain/shared"
)

func (s *service) ListOrders(ctx context.Context, page, limit int) (*ListOrdersResult, error) {
	p := shared.NewPagination(page, limit)
	items, err := s.orders.List(ctx, p)
	if err != nil {
		return nil, err
	}
	return &ListOrdersResult{Items: items, Page: p.Page, Limit: p.Limit}, nil
}
