package category

import (
	"context"
	"goCartAPI/domain/shared"
)

type Repository interface {
	Save(ctx context.Context, category *Category) error
	FindByID(ctx context.Context, id shared.ID) (*Category, error)
	FindBySlug(ctx context.Context, slug string) (*Category, error)
	List(ctx context.Context, pagination shared.Pagination) ([]*Category, error)
	Count(ctx context.Context) (int, error)
	Delete(ctx context.Context, id shared.ID) error
}
