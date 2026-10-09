package order

func CanTransition(from, to Status) bool {

	transitions := map[Status]map[Status]bool{

		StatusPending: {
			StatusAwaitingPayment: true,
			StatusCancelled:       true,
		},

		StatusAwaitingPayment: {
			StatusPaid:      true,
			StatusCancelled: true,
		},

		StatusPaid: {
			StatusProcessing: true,
			StatusRefunded:   true,
		},

		StatusProcessing: {
			StatusShipped:   true,
			StatusCancelled: true,
		},
		StatusShipped: {
			StatusDelivered: true,
		},
		StatusDelivered: {
			StatusRefunded: true,
		},
	}
	return transitions[from][to]
}

func CanCancel(status Status) bool {
	return status == StatusPending || status == StatusAwaitingPayment || status == StatusProcessing
}
