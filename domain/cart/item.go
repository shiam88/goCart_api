package cart

import (
	"goCartAPI/domain/shared"
	"strings"
)

type Item struct {
	productID   shared.ID
	productName string
	unitPrice   shared.Money
	quantity    int
}

func NewItem(productID shared.ID, productName string, unitPrice shared.Money, quantity int) (*Item, error) {
	if shared.IsZeroID(productID) {
		return nil, ErrInvalidProductID
	}
	productName = strings.TrimSpace(productName)
	if productName == "" {
		return nil, ErrEmptyProductName
	}
	if quantity <= 0 {
		return nil, ErrInvalidQuantity
	}
	return &Item{productID: productID, productName: productName, unitPrice: unitPrice, quantity: quantity}, nil
}

func (i *Item) ProductID() shared.ID {
	return i.productID
}

func (i *Item) ProductName() string {
	return i.productName
}
func (i *Item) UnitPrice() shared.Money {
	return i.unitPrice
}
func (i *Item) Quantity() int {
	return i.quantity
}

func (i *Item) UpdateQuantity(quantity int) error {
	if quantity <= 0 {
		return ErrInvalidQuantity
	}
	i.quantity = quantity
	return nil
}

func (i *Item) Subtotal() (shared.Money, error) {
	return i.unitPrice.Multiply(i.quantity)
}
