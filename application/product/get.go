package product

import (
	"context"
	"goCartAPI/domain/product"
	"goCartAPI/domain/shared"
)

func (s *service) Get(ctx context.Context, id shared.ID) (*product.Product, error) {
	return s.products.FindByID(ctx, id)
}
