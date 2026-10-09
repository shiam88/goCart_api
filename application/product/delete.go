package product

import (
	"context"
	"goCartAPI/domain/shared"
)

func (s *service) Delete(ctx context.Context, id shared.ID) error {
	return s.products.Delete(ctx, id)
}
