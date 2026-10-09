package payment

import (
	"context"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
)

type Service interface {
	Initiate(ctx context.Context, orderID shared.ID) (string, error)
	Get(ctx context.Context, id shared.ID) (*payment.Transaction, error)
	Confirm(ctx context.Context, orderID shared.ID, providerReference string, success bool) error
	Refund(ctx context.Context, transactionID shared.ID) error
}

type TransactionManager interface {
	WithinTransaction(ctx context.Context, fn func(context.Context) error) error
}

type PaymentGatewayResolver interface {
	Resolve(provider payment.Provider) (payment.Gateway, error)
}
