package cart

import "errors"

var (
	ErrInvalidCartID    = errors.New("invalid cart id")
	ErrInvalidUserID    = errors.New("invalid user id")
	ErrInvalidProductID = errors.New("invalid product id")
	ErrInvalidQuantity  = errors.New("quantity must be greater than zero")
	ErrItemNotFound     = errors.New("cart item not found")
	ErrEmptyProductName = errors.New("product name is required")
	ErrEmptyCart        = errors.New("cart is empty")
	ErrCartNotFound     = errors.New("cart not found")
	ErrCurrencyMismatch = errors.New("Currency not match")
)
