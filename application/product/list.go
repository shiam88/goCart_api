package product

import (
	"context"
	"goCartAPI/domain/product"
	"goCartAPI/domain/shared"
)

func (s *service) List(ctx context.Context, in ListInput) (*ListResult, error) {
	p := shared.NewPagination(in.Page, in.Limit)
	f := product.Filter{Search: in.Search, CategoryID: in.CategoryID, MinPrice: in.MinPrice, MaxPrice: in.MaxPrice, OnlyActive: in.OnlyActive}
	items, err := s.products.List(ctx, f, p)
	if err != nil {
		return nil, err
	}
	total, err := s.products.Count(ctx, f)
	if err != nil {
		return nil, err
	}
	return &ListResult{Items: items, Total: total, Page: p.Page, Limit: p.Limit}, nil
}
