package paymentinfra

import (
	"fmt"

	"goCartAPI/domain/payment"
)

type ProviderFactory struct {
	mock payment.Gateway
	cod  payment.Gateway
}

func NewProviderFactory(mock, cod payment.Gateway) *ProviderFactory {
	return &ProviderFactory{mock: mock, cod: cod}
}

func (f *ProviderFactory) Resolve(provider payment.Provider) (payment.Gateway, error) {
	switch provider {
	case payment.ProviderMock:
		return f.mock, nil
	case payment.ProviderCOD:
		return f.cod, nil
	default:
		return nil, fmt.Errorf("payment provider %s is not configured in v1", provider)
	}
}
