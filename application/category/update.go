package category

import (
	"context"
	"goCartAPI/domain/category"
	"goCartAPI/domain/shared"
)

func (s *service) Update(ctx context.Context, id shared.ID, name string, active *bool) (*category.Category, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if name != "" {
		if err := c.Rename(name); err != nil {
			return nil, err
		}
	}
	if active != nil {
		if *active {
			c.Activate()
		} else {
			c.Deactivate()
		}
	}
	if err := s.repo.Save(ctx, c); err != nil {
		return nil, err
	}
	return c, nil
}
