package order

import (
	"goCartAPI/domain/shared"
	"strings"
	"time"
)

type Order struct {
	id              shared.ID
	userID          shared.ID
	items           []Item
	status          Status
	subtotal        shared.Money
	shippingFee     shared.Money
	discount        shared.Money
	total           shared.Money
	shippingAddress string
	createdAt       time.Time
	updatedAt       time.Time
}

func NewOrder(userID shared.ID, items []Item, shippingFee, discount shared.Money, shippingAddress string) (*Order, error) {
	if shared.IsZeroID(userID) {
		return nil, ErrInvalidUserID
	}
	if len(items) == 0 {
		return nil, ErrEmptyOrder
	}
	shippingAddress = strings.TrimSpace(shippingAddress)
	if shippingAddress == "" {
		return nil, ErrShippingAddress
	}
	currency := items[0].UnitPrice().Currency()
	if shippingFee.Currency() != currency || discount.Currency() != currency {
		return nil, shared.ErrCurrencyMismatch
	}
	subtotal, err := shared.ZeroMoney(currency)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.UnitPrice().Currency() != currency {
			return nil, shared.ErrCurrencyMismatch
		}
		line, err := item.Subtotal()
		if err != nil {
			return nil, err
		}
		subtotal, err = subtotal.Add(line)
		if err != nil {
			return nil, err
		}
	}
	beforeDiscount, err := subtotal.Add(shippingFee)
	if err != nil {
		return nil, err
	}
	total, err := beforeDiscount.Subtract(discount)
	if err != nil {
		return nil, err
	}
	now := time.Now().UTC()
	return &Order{id: shared.NewID(), userID: userID, items: append([]Item(nil), items...), status: StatusPending, subtotal: subtotal, shippingFee: shippingFee, discount: discount, total: total, shippingAddress: shippingAddress, createdAt: now, updatedAt: now}, nil
}

func Restore(id, userID shared.ID, items []Item, status Status, subtotal, shippingFee, discount, total shared.Money, shippingAddress string, createdAt, updatedAt time.Time) (*Order, error) {
	if shared.IsZeroID(id) {
		return nil, ErrInvalidOrderID
	}
	if shared.IsZeroID(userID) {
		return nil, ErrInvalidUserID
	}
	if len(items) == 0 {
		return nil, ErrEmptyOrder
	}
	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}
	if strings.TrimSpace(shippingAddress) == "" {
		return nil, ErrShippingAddress
	}
	currency := subtotal.Currency()
	if shippingFee.Currency() != currency || discount.Currency() != currency || total.Currency() != currency {
		return nil, shared.ErrCurrencyMismatch
	}
	return &Order{id: id, userID: userID, items: append([]Item(nil), items...), status: status, subtotal: subtotal, shippingFee: shippingFee, discount: discount, total: total, shippingAddress: shippingAddress, createdAt: createdAt.UTC(), updatedAt: updatedAt.UTC()}, nil
}

func (o *Order) ID() shared.ID {
	return o.id
}
func (o *Order) UserID() shared.ID {
	return o.userID
}

func (o *Order) Items() []Item {
	return append([]Item(nil), o.items...)
}

func (o *Order) Status() Status {
	return o.status
}
func (o *Order) Subtotal() shared.Money {
	return o.subtotal
}
func (o *Order) ShippingFee() shared.Money {
	return o.shippingFee
}
func (o *Order) Discount() shared.Money {
	return o.discount
}

func (o *Order) Total() shared.Money {
	return o.total
}

func (o *Order) ShippingAddress() string {
	return o.shippingAddress
}

func (o *Order) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Order) UpdatedAt() time.Time {
	return o.updatedAt
}

func (o *Order) TransitionTo(next Status) error {
	if !next.IsValid() {
		return ErrInvalidStatus
	}
	if !CanTransition(o.status, next) {
		return ErrInvalidTransition
	}
	o.status = next
	o.updatedAt = time.Now().UTC()
	return nil
}
func (o *Order) Cancel() error {
	if !CanCancel(o.status) {
		return ErrOrderCannotCancel
	}
	o.status = StatusCancelled
	o.updatedAt = time.Now().UTC()
	return nil
}
