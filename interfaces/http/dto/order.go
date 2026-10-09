package dto

type CheckoutItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type CheckoutRequest struct {
	Source          string                `json:"source"`
	Items           []CheckoutItemRequest `json:"items,omitempty"`
	ShippingAddress string                `json:"shipping_address"`
	PaymentMethod   string                `json:"payment_method"`
}
type OrderItemResponse struct {
	ProductID      string `json:"product_id"`
	ProductName    string `json:"product_name"`
	UnitPriceMinor int64  `json:"unit_price_minor"`
	Currency       string `json:"currency"`
	Quantity       int    `json:"quantity"`
}
type OrderResponse struct {
	ID               string              `json:"id"`
	UserID           string              `json:"user_id"`
	Status           string              `json:"status"`
	Items            []OrderItemResponse `json:"items"`
	SubtotalMinor    int64               `json:"subtotal_minor"`
	ShippingFeeMinor int64               `json:"shipping_fee_minor"`
	DiscountMinor    int64               `json:"discount_minor"`
	TotalMinor       int64               `json:"total_minor"`
	Currency         string              `json:"currency"`
	ShippingAddress  string              `json:"shipping_address"`
}
type CheckoutResponse struct {
	Order             OrderResponse   `json:"order"`
	Payment           PaymentResponse `json:"payment"`
	ProviderReference string          `json:"provider_reference,omitempty"`
}
type UpdateOrderStatusRequest struct {
	Status string `json:"status"`
}
