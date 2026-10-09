package product

import (
	"context"

	"goCartAPI/domain/product"
	"goCartAPI/domain/shared"
)

type Service interface {
	Create(ctx context.Context, in CreateInput) (*product.Product, error)
	Get(ctx context.Context, id shared.ID) (*product.Product, error)
	List(ctx context.Context, in ListInput) (*ListResult, error)
	Update(ctx context.Context, in UpdateInput) (*product.Product, error)
	Delete(ctx context.Context, id shared.ID) error
}

type TransactionManager interface {
	WithinTransaction(ctx context.Context, fn func(context.Context) error) error
}

type CreateInput struct {
	CategoryID   shared.ID
	Name         string
	SKU          string
	Description  string
	Price        shared.Money
	ImageURL     string
	InitialStock int
}

type UpdateInput struct {
	ID                          shared.ID
	Name, Description, ImageURL string
	CategoryID                  *shared.ID
	Price                       *shared.Money
	Active                      *bool
}

type ListInput struct {
	Page, Limit        int
	Search             string
	CategoryID         *shared.ID
	MinPrice, MaxPrice *int64
	OnlyActive         bool
}

type ListResult struct {
	Items              []*product.Product
	Total, Page, Limit int
}
