package cart

import (
	"goCartAPI/domain/cart"
	"goCartAPI/domain/product"
)

type service struct {
	carts    cart.Repository
	products product.Repository
	currency string
}

func NewService(carts cart.Repository, products product.Repository, currency string) Service {
	return &service{carts: carts, products: products, currency: currency}
}
