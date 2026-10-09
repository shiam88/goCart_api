package category

import (
	"context"
	"goCartAPI/domain/shared"
)

func (s *service) Delete(ctx context.Context, id shared.ID) error {
	return s.repo.Delete(ctx, id)
}
