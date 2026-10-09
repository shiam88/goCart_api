package product

import (
	"goCartAPI/domain/category"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/product"
)

type service struct {
	products   product.Repository
	categories category.Repository
	inventory  inventory.Repository
	tx         TransactionManager
}

func NewService(products product.Repository, categories category.Repository, inventoryRepo inventory.Repository, tx TransactionManager) Service {
	return &service{products: products, categories: categories, inventory: inventoryRepo, tx: tx}
}
