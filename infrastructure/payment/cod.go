package paymentinfra

import (
	"context"
	"fmt"

	"goCartAPI/domain/shared"

	"github.com/google/uuid"
)

type CODGateway struct{}

func NewCODGateway() *CODGateway {
	return &CODGateway{}
}

func (g *CODGateway) Initiate(ctx context.Context, orderID shared.ID, amount shared.Money) (string, error) {
	return "COD-" + uuid.NewString(), nil
}

func (g *CODGateway) Refund(ctx context.Context, ref string, amount shared.Money) error {
	return fmt.Errorf("COD refund must be handled by business operations, not an online gateway")
}
