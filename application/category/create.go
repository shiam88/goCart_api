package category

import (
	"context"
	"goCartAPI/domain/category"
)

func (s *service) Create(ctx context.Context, name string) (*category.Category, error) {
	c, err := category.NewCategory(name)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
