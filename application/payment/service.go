package payment

import (
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
)

type service struct {
	payments  payment.Repository
	orders    order.Repository
	inventory inventory.Repository
	gateways  PaymentGatewayResolver
	tx        TransactionManager
}

func NewService(payments payment.Repository, orders order.Repository, inventoryRepo inventory.Repository, gateways PaymentGatewayResolver, tx TransactionManager) Service {
	return &service{payments: payments, orders: orders, inventory: inventoryRepo, gateways: gateways, tx: tx}
}
