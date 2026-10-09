package order

import "errors"

var (
	ErrInvalidOrderID    = errors.New("invalid order id")
	ErrInvalidUserID     = errors.New("invalid user id")
	ErrInvalidProductID  = errors.New("invalid product id")
	ErrInvalidQuantity   = errors.New("quantity must be greater than zero")
	ErrEmptyProductName  = errors.New("product name is required")
	ErrEmptyOrder        = errors.New("order must contain at least one item")
	ErrInvalidStatus     = errors.New("invalid order status")
	ErrInvalidTransition = errors.New("invalid order status transition")
	ErrOrderCannotCancel = errors.New("order cannot be cancelled in its current state")
	ErrOrderNotFound     = errors.New("order not found")
	ErrShippingAddress   = errors.New("shipping address is required")
)
