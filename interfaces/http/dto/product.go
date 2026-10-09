package dto

type CreateProductRequest struct {
	CategoryID   string `json:"category_id"`
	Name         string `json:"name"`
	SKU          string `json:"sku"`
	Description  string `json:"description"`
	PriceMinor   int64  `json:"price_minor"`
	Currency     string `json:"currency"`
	ImageURL     string `json:"image_url"`
	InitialStock int    `json:"initial_stock"`
}
type UpdateProductRequest struct {
	Name        string `json:"name,omitempty"`
	CategoryID  string `json:"category_id,omitempty"`
	Description string `json:"description,omitempty"`
	PriceMinor  *int64 `json:"price_minor,omitempty"`
	Currency    string `json:"currency,omitempty"`
	ImageURL    string `json:"image_url,omitempty"`
	Active      *bool  `json:"active,omitempty"`
}
type ProductResponse struct {
	ID          string `json:"id"`
	CategoryID  string `json:"category_id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	SKU         string `json:"sku"`
	Description string `json:"description"`
	PriceMinor  int64  `json:"price_minor"`
	Currency    string `json:"currency"`
	ImageURL    string `json:"image_url"`
	Active      bool   `json:"active"`
}
