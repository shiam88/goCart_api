package order

type Status string

const (
	StatusPending         Status = "PENDING"
	StatusAwaitingPayment Status = "AWAITING_PAYMENT"
	StatusPaid            Status = "PAID"
	StatusProcessing      Status = "PROCESSING"
	StatusShipped         Status = "SHIPPED"
	StatusDelivered       Status = "DELIVERED"
	StatusCancelled       Status = "CANCELLED"
	StatusRefunded        Status = "REFUNDED"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending,
		StatusAwaitingPayment,
		StatusPaid,
		StatusProcessing,
		StatusShipped,
		StatusDelivered,
		StatusCancelled,
		StatusRefunded:
		return true

	default:
		return false
	}
}
