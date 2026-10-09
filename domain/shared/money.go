package shared

import (
	"math"
	"strings"
)

// Money is a value object. Amount is stored in the smallest currency unit
// (for example poisha/cents) to avoid floating-point rounding errors.
type Money struct {
	amount   int64
	currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
	if amount < 0 {
		return Money{}, ErrInvalidAmount
	}
	currency = strings.ToUpper(strings.TrimSpace(currency))
	if len(currency) != 3 {
		return Money{}, ErrInvalidCurrency
	}
	return Money{amount: amount, currency: currency}, nil
}

func ZeroMoney(currency string) (Money, error) { return NewMoney(0, currency) }
func (m Money) Amount() int64                  { return m.amount }
func (m Money) Currency() string               { return m.currency }
func (m Money) IsZero() bool                   { return m.amount == 0 }

func (m Money) Equal(other Money) bool {
	return m.amount == other.amount && m.currency == other.currency
}

func (m Money) Add(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, ErrCurrencyMismatch
	}
	if other.amount > 0 && m.amount > math.MaxInt64-other.amount {
		return Money{}, ErrInvalidAmount
	}
	return Money{amount: m.amount + other.amount, currency: m.currency}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.currency != other.currency {
		return Money{}, ErrCurrencyMismatch
	}
	if other.amount > m.amount {
		return Money{}, ErrInsufficientMoney
	}
	return Money{amount: m.amount - other.amount, currency: m.currency}, nil
}

func (m Money) Multiply(quantity int) (Money, error) {
	if quantity < 0 {
		return Money{}, ErrInvalidMultiplier
	}
	if quantity == 0 || m.amount == 0 {
		return Money{amount: 0, currency: m.currency}, nil
	}
	if m.amount > math.MaxInt64/int64(quantity) {
		return Money{}, ErrInvalidAmount
	}
	return Money{amount: m.amount * int64(quantity), currency: m.currency}, nil
}
