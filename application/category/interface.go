package category

import (
	"context"
	"goCartAPI/domain/category"
	"goCartAPI/domain/shared"
)

type Service interface {
	Create(ctx context.Context, name string) (*category.Category, error)
	List(ctx context.Context, page, limit int) (*ListResult, error)
	Update(ctx context.Context, id shared.ID, name string, active *bool) (*category.Category, error)
	Delete(ctx context.Context, id shared.ID) error
}

type ListResult struct {
	Items []*category.Category
	Total int
	Page  int
	Limit int
}
