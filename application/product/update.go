package product

import (
	"context"
	"goCartAPI/domain/product"
)

func (s *service) Update(ctx context.Context, in UpdateInput) (*product.Product, error) {
	p, err := s.products.FindByID(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	if in.Name != "" {
		if err := p.Rename(in.Name); err != nil {
			return nil, err
		}
	}
	if in.CategoryID != nil {
		if err := p.ChangeCategory(*in.CategoryID); err != nil {
			return nil, err
		}
	}
	if in.Price != nil {
		if err := p.ChangePrice(*in.Price); err != nil {
			return nil, err
		}
	}
	if in.Description != "" || in.ImageURL != "" {
		p.UpdateDetails(in.Description, in.ImageURL)
	}
	if in.Active != nil {
		if *in.Active {
			p.Activate()
		} else {
			p.Deactivate()
		}
	}
	if err := s.products.Save(ctx, p); err != nil {
		return nil, err
	}
	return p, nil
}
