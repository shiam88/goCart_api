package payment

type Status string

const (
	StatusPending    Status = "PENDING"
	StatusProcessing Status = "PROCESSING"
	StatusPaid       Status = "PAID"
	StatusFailed     Status = "FAILED"
	StatusCancelled  Status = "CANCELLED"
	StatusRefunded   Status = "REFUNDED"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending,
		StatusProcessing,
		StatusPaid,
		StatusFailed,
		StatusCancelled,
		StatusRefunded:

		return true

	default:
		return false
	}
}
