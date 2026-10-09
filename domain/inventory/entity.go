package inventory

import (
	"goCartAPI/domain/shared"
	"time"
)

type Inventory struct {
	productID shared.ID
	onHand    int
	reserved  int
	updatedAt time.Time
}

func NewInventory(productID shared.ID, initialStock int) (*Inventory, error) {
	if shared.IsZeroID(productID) {
		return nil, ErrInvalidProductID
	}
	if initialStock < 0 {
		return nil, ErrInvalidQuantity
	}
	return &Inventory{productID: productID, onHand: initialStock, updatedAt: time.Now().UTC()}, nil
}

func Restore(productID shared.ID, onHand, reserved int, updatedAt time.Time) (*Inventory, error) {
	if shared.IsZeroID(productID) {
		return nil, ErrInvalidProductID
	}
	if onHand < 0 || reserved < 0 || reserved > onHand {
		return nil, ErrInvalidQuantity
	}
	return &Inventory{productID: productID, onHand: onHand, reserved: reserved, updatedAt: updatedAt.UTC()}, nil
}

func (i *Inventory) ProductID() shared.ID      { return i.productID }
func (i *Inventory) OnHand() int               { return i.onHand }
func (i *Inventory) Reserved() int             { return i.reserved }
func (i *Inventory) Available() int            { return i.onHand - i.reserved }
func (i *Inventory) UpdatedAt() time.Time      { return i.updatedAt }
func (i *Inventory) HasAvailable(qty int) bool { return qty > 0 && i.Available() >= qty }

func (i *Inventory) AddStock(qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}
	i.onHand += qty
	i.touch()
	return nil
}

func (i *Inventory) Reserve(qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}
	if i.Available() < qty {
		return ErrInsufficientStock
	}
	i.reserved += qty
	i.touch()
	return nil
}

func (i *Inventory) Release(qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}
	if i.reserved < qty {
		return ErrInsufficientReserve
	}
	i.reserved -= qty
	i.touch()
	return nil
}

// CommitReservation finalizes a sale: reserved quantity leaves physical stock.
func (i *Inventory) CommitReservation(qty int) error {
	if qty <= 0 {
		return ErrInvalidQuantity
	}
	if i.reserved < qty {
		return ErrInsufficientReserve
	}
	i.reserved -= qty
	i.onHand -= qty
	i.touch()
	return nil
}

func (i *Inventory) touch() { i.updatedAt = time.Now().UTC() }
