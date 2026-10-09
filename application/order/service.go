package order

import (
	"goCartAPI/domain/cart"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/product"
	"goCartAPI/domain/shared"
)

type service struct {
	products    product.Repository
	inventory   inventory.Repository
	carts       cart.Repository
	orders      order.Repository
	payments    payment.Repository
	tx          TransactionManager
	gateways    PaymentGatewayResolver
	shippingFee shared.Money
}

func NewService(products product.Repository, inventoryRepo inventory.Repository, carts cart.Repository, orders order.Repository, payments payment.Repository, tx TransactionManager, gateways PaymentGatewayResolver, shippingFee shared.Money) Service {
	return &service{products: products, inventory: inventoryRepo, carts: carts, orders: orders, payments: payments, tx: tx, gateways: gateways, shippingFee: shippingFee}
}
