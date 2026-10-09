package payment

import (
	"context"
	"goCartAPI/domain/shared"
)

type Provider string

const (
	ProviderCOD    Provider = "COD"
	ProviderMock   Provider = "MOCK"
	ProviderBKash  Provider = "BKASH"
	ProviderNagad  Provider = "NAGAD"
	ProviderStripe Provider = "STRIPE"
)

func (p Provider) IsValid() bool {
	switch p {
	case ProviderCOD,
		ProviderMock,
		ProviderBKash,
		ProviderNagad,
		ProviderStripe:

		return true

	default:
		return false
	}
}

// gateway interface
type Gateway interface {
	Initiate(
		ctx context.Context,
		orderID shared.ID,
		amount shared.Money,
	) (string, error)

	Refund(
		ctx context.Context,
		providerReference string,
		amount shared.Money,
	) error
}
