package order

import (
	"context"
	"goCartAPI/domain/shared"
)

func (s *service) List(ctx context.Context, userID shared.ID, page, limit int) (*ListResult, error) {
	p := shared.NewPagination(page, limit)
	items, err := s.orders.ListByUserID(ctx, userID, p)
	if err != nil {
		return nil, err
	}
	total, err := s.orders.CountByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return &ListResult{Items: items, Total: total, Page: p.Page, Limit: p.Limit}, nil
}
