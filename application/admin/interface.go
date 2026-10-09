package admin

import (
	"context"

	"goCartAPI/domain/inventory"

	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
)

type Service interface {
	ListOrders(ctx context.Context, page, limit int) (*ListOrdersResult, error)
	UpdateOrderStatus(ctx context.Context, id shared.ID, status order.Status) (*order.Order, error)
	AdjustStock(ctx context.Context, productID shared.ID, delta int) (*inventory.Inventory, error)
}

type OrderService interface {
	UpdateStatus(ctx context.Context, orderID shared.ID, target order.Status) (*order.Order, error)
}

type InventoryService interface {
	Adjust(ctx context.Context, productID shared.ID, delta int) (*inventory.Inventory, error)
}

type ListOrdersResult struct {
	Items       []*order.Order
	Page, Limit int
}
