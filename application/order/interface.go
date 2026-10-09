package order

import (
	"context"
	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
)

type Service interface {
	Checkout(ctx context.Context, in CheckoutInput) (*CheckoutResult, error)
	Get(ctx context.Context, id, userID shared.ID, admin bool) (*order.Order, error)
	List(ctx context.Context, userID shared.ID, page, limit int) (*ListResult, error)
	Cancel(ctx context.Context, orderID, userID shared.ID) error
	UpdateStatus(ctx context.Context, orderID shared.ID, target order.Status) (*order.Order, error)
}

type TransactionManager interface {
	WithinTransaction(ctx context.Context, fn func(context.Context) error) error
}

type PaymentGatewayResolver interface {
	Resolve(provider payment.Provider) (payment.Gateway, error)
}

type CheckoutItemInput struct {
	ProductID shared.ID
	Quantity  int
}

type CheckoutInput struct {
	UserID          shared.ID
	Source          string
	Items           []CheckoutItemInput
	ShippingAddress string
	PaymentProvider payment.Provider
}

type CheckoutResult struct {
	Order             *order.Order
	Payment           *payment.Transaction
	ProviderReference string
}

type ListResult struct {
	Items              []*order.Order
	Total, Page, Limit int
}
