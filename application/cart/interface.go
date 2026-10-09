package cart

import (
	"context"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/shared"
)

type Service interface {
	Get(ctx context.Context, userID shared.ID) (*cart.Cart, error)
	AddItem(ctx context.Context, userID, productID shared.ID, qty int) (*cart.Cart, error)
	UpdateItem(ctx context.Context, userID, productID shared.ID, qty int) (*cart.Cart, error)
	RemoveItem(ctx context.Context, userID, productID shared.ID) (*cart.Cart, error)
	Clear(ctx context.Context, userID shared.ID) (*cart.Cart, error)
}
