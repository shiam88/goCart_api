package product

import (
	"context"

	"goCartAPI/domain/shared"
)

type Filter struct {
	Search     string
	CategoryID *shared.ID
	MinPrice   *int64
	MaxPrice   *int64
	OnlyActive bool
}

type Repository interface {
	Save(ctx context.Context, product *Product) error
	FindByID(ctx context.Context, id shared.ID) (*Product, error)
	FindBySKU(ctx context.Context, sku string) (*Product, error)
	List(ctx context.Context, filter Filter, pagination shared.Pagination) ([]*Product, error)
	Count(ctx context.Context, filter Filter) (int, error)
	Delete(ctx context.Context, id shared.ID) error
	SKUExists(ctx context.Context, sku string) (bool, error)
}
