package admin

import "goCartAPI/domain/order"

type service struct {
	orders   order.Repository
	orderSvc OrderService
	stockSvc InventoryService
}

func NewService(orders order.Repository, orderSvc OrderService, stockSvc InventoryService) Service {
	return &service{
		orders:   orders,
		orderSvc: orderSvc,
		stockSvc: stockSvc,
	}
}
