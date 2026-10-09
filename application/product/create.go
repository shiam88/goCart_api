package product

import (
	"context"

	"goCartAPI/domain/inventory"
	"goCartAPI/domain/product"
)

func (s *service) Create(ctx context.Context, in CreateInput) (*product.Product, error) {
	var created *product.Product
	err := s.tx.WithinTransaction(ctx, func(txctx context.Context) error {
		if _, err := s.categories.FindByID(txctx, in.CategoryID); err != nil {
			return err
		}
		exists, err := s.products.SKUExists(txctx, in.SKU)
		if err != nil {
			return err
		}
		if exists {
			return product.ErrSKUAlreadyExists
		}
		p, err := product.NewProduct(in.CategoryID, in.Name, in.SKU, in.Description, in.Price, in.ImageURL)
		if err != nil {
			return err
		}
		inv, err := inventory.NewInventory(p.ID(), in.InitialStock)
		if err != nil {
			return err
		}
		if err := s.products.Save(txctx, p); err != nil {
			return err
		}
		if err := s.inventory.Save(txctx, inv); err != nil {
			return err
		}
		created = p
		return nil
	})
	return created, err
}
