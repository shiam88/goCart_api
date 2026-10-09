package category

import (
	"context"
	"goCartAPI/domain/shared"
)

func (s *service) List(ctx context.Context, page, limit int) (*ListResult, error) {
	p := shared.NewPagination(page, limit)
	items, err := s.repo.List(ctx, p)
	if err != nil {
		return nil, err
	}
	total, err := s.repo.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &ListResult{Items: items, Total: total, Page: p.Page, Limit: p.Limit}, nil
}
