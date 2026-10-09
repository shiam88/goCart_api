package payment

import "errors"

var (
	ErrInvalidTransactionID = errors.New("invalid transaction id")
	ErrInvalidOrderID       = errors.New("invalid order id")
	ErrInvalidUserID        = errors.New("invalid user id")
	ErrInvalidAmount        = errors.New("payment amount must be greater than zero")
	ErrInvalidProvider      = errors.New("invalid payment provider")
	ErrInvalidStatus        = errors.New("invalid payment status")
	ErrInvalidTransition    = errors.New("invalid payment status transition")
	ErrTransactionNotFound  = errors.New("payment transaction not found")
)
