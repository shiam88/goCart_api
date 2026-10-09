package shared

import "errors"

var (
	ErrInvalidID         = errors.New("invalid id")
	ErrInvalidAmount     = errors.New("amount cannot be negative")
	ErrInvalidCurrency   = errors.New("currency must be a three-letter code")
	ErrCurrencyMismatch  = errors.New("money currencies do not match")
	ErrInsufficientMoney = errors.New("insufficient money")
	ErrInvalidMultiplier = errors.New("multiplier cannot be negative")
	ErrInvalidPagination = errors.New("invalid pagination values")
)
