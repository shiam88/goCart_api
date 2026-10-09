package inventory

import "errors"

var (
	ErrInvalidProductID    = errors.New("invalid product id")
	ErrInvalidQuantity     = errors.New("invalid quantity")
	ErrInsufficientStock   = errors.New("insufficient stock")
	ErrInsufficientReserve = errors.New("insufficient reserved stock")
	ErrInventoryNotFound   = errors.New("inventory not found")
)
