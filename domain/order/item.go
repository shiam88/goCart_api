package order

import (
	"strings"

	"goCartAPI/domain/shared"
)

type Item struct {
	productID   shared.ID
	productName string
	unitPrice   shared.Money
	quantity    int
}

func NewItem(
	productID shared.ID,
	productName string,
	unitPrice shared.Money,
	quantity int,
) (Item, error) {

	if shared.IsZeroID(productID) {
		return Item{}, ErrInvalidProductID
	}

	if strings.TrimSpace(productName) == "" {
		return Item{}, ErrEmptyProductName
	}

	if quantity <= 0 {
		return Item{}, ErrInvalidQuantity
	}

	return Item{
		productID:   productID,
		productName: productName,
		unitPrice:   unitPrice,
		quantity:    quantity,
	}, nil
}

func (i Item) ProductID() shared.ID {
	return i.productID
}

func (i Item) ProductName() string {
	return i.productName
}

func (i Item) UnitPrice() shared.Money {
	return i.unitPrice
}

func (i Item) Quantity() int {
	return i.quantity
}

func (i Item) Subtotal() (shared.Money, error) {
	return i.unitPrice.Multiply(i.quantity)
}
