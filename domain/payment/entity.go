package payment

import (
	"goCartAPI/domain/shared"
	"strings"
	"time"
)

type Transaction struct {
	id                shared.ID
	orderID           shared.ID
	userID            shared.ID
	amount            shared.Money
	provider          Provider
	status            Status
	providerReference string
	createdAt         time.Time
	updatedAt         time.Time
}

func NewTransaction(orderID, userID shared.ID, amount shared.Money, provider Provider) (*Transaction, error) {
	if shared.IsZeroID(orderID) {
		return nil, ErrInvalidOrderID
	}
	if shared.IsZeroID(userID) {
		return nil, ErrInvalidUserID
	}
	if amount.Amount() <= 0 {
		return nil, ErrInvalidAmount
	}
	if !provider.IsValid() {
		return nil, ErrInvalidProvider
	}
	now := time.Now().UTC()
	return &Transaction{id: shared.NewID(), orderID: orderID, userID: userID, amount: amount, provider: provider, status: StatusPending, createdAt: now, updatedAt: now}, nil
}

func Restore(id, orderID, userID shared.ID, amount shared.Money, provider Provider, status Status, providerReference string, createdAt, updatedAt time.Time) (*Transaction, error) {
	if shared.IsZeroID(id) {
		return nil, ErrInvalidTransactionID
	}
	t, err := NewTransaction(orderID, userID, amount, provider)
	if err != nil {
		return nil, err
	}
	if !status.IsValid() {
		return nil, ErrInvalidStatus
	}
	t.id = id
	t.status = status
	t.providerReference = strings.TrimSpace(providerReference)
	t.createdAt = createdAt.UTC()
	t.updatedAt = updatedAt.UTC()
	return t, nil
}

func (t *Transaction) ID() shared.ID {
	return t.id
}

func (t *Transaction) OrderID() shared.ID {
	return t.orderID
}

func (t *Transaction) UserID() shared.ID {
	return t.userID
}

func (t *Transaction) Amount() shared.Money {
	return t.amount
}

func (t *Transaction) Provider() Provider {
	return t.provider
}

func (t *Transaction) Status() Status {
	return t.status
}

func (t *Transaction) ProviderReference() string {
	return t.providerReference
}

func (t *Transaction) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Transaction) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Transaction) StartProcessing(providerReference string) error {
	if t.status != StatusPending {
		return ErrInvalidTransition
	}
	t.status = StatusProcessing
	t.providerReference = strings.TrimSpace(providerReference)
	t.touch()
	return nil
}
func (t *Transaction) MarkPaid(providerReference string) error {
	if t.status != StatusPending && t.status != StatusProcessing {
		return ErrInvalidTransition
	}
	t.status = StatusPaid
	if strings.TrimSpace(providerReference) != "" {
		t.providerReference = strings.TrimSpace(providerReference)
	}
	t.touch()
	return nil
}
func (t *Transaction) MarkFailed() error {
	if t.status != StatusPending && t.status != StatusProcessing {
		return ErrInvalidTransition
	}
	t.status = StatusFailed
	t.touch()
	return nil
}
func (t *Transaction) Cancel() error {
	if t.status != StatusPending {
		return ErrInvalidTransition
	}
	t.status = StatusCancelled
	t.touch()
	return nil
}
func (t *Transaction) Refund() error {
	if t.status != StatusPaid {
		return ErrInvalidTransition
	}
	t.status = StatusRefunded
	t.touch()
	return nil
}

func (t *Transaction) touch() {
	t.updatedAt = time.Now().UTC()
}
