package paymentinfra

import (
	"context"
	"fmt"

	"goCartAPI/domain/shared"

	"github.com/google/uuid"
)

type MockGateway struct{}

func NewMockGateway() *MockGateway {
	return &MockGateway{}
}

func (g *MockGateway) Initiate(ctx context.Context, orderID shared.ID, amount shared.Money) (string, error) {
	return "MOCK-" + uuid.NewString(), nil
}

func (g *MockGateway) Refund(ctx context.Context, ref string, amount shared.Money) error {
	if ref == "" {
		return fmt.Errorf("provider reference is required")
	}
	return nil
}
