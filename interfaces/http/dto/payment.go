package dto

type PaymentResponse struct {
	ID                string `json:"id"`
	OrderID           string `json:"order_id"`
	Provider          string `json:"provider"`
	Status            string `json:"status"`
	AmountMinor       int64  `json:"amount_minor"`
	Currency          string `json:"currency"`
	ProviderReference string `json:"provider_reference,omitempty"`
}
type PaymentCallbackRequest struct {
	OrderID           string `json:"order_id"`
	ProviderReference string `json:"provider_reference"`
	Success           bool   `json:"success"`
}
