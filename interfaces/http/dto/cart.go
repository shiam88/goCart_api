package dto

type AddCartItemRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
type UpdateCartItemRequest struct {
	Quantity int `json:"quantity"`
}
type CartItemResponse struct {
	ProductID      string `json:"product_id"`
	ProductName    string `json:"product_name"`
	UnitPriceMinor int64  `json:"unit_price_minor"`
	Currency       string `json:"currency"`
	Quantity       int    `json:"quantity"`
	SubtotalMinor  int64  `json:"subtotal_minor"`
}
type CartResponse struct {
	ID            string             `json:"id"`
	Currency      string             `json:"currency"`
	Items         []CartItemResponse `json:"items"`
	TotalMinor    int64              `json:"total_minor"`
	TotalQuantity int                `json:"total_quantity"`
}
