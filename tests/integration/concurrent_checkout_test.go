package integration

import "testing"

func TestConcurrentCheckout(t *testing.T) {
	t.Skip("concurrent stock protection requires a disposable PostgreSQL integration environment")
}
