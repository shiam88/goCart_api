package dto

import (
	"goCartAPI/domain/cart"
	"goCartAPI/domain/category"

	"goCartAPI/domain/inventory"
	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/product"
)

type PaginationResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items,omitempty"`
	TotalPages int `json:"total_pages,omitempty"`
}
type InventoryResponse struct {
	ProductID                   string `json:"product_id"`
	OnHand, Reserved, Available int
}

func CategoryFromDomain(c *category.Category) CategoryResponse {
	return CategoryResponse{ID: c.ID().String(), Name: c.Name(), Slug: c.Slug(), Active: c.IsActive()}
}

func ProductFromDomain(p *product.Product) ProductResponse {
	return ProductResponse{ID: p.ID().String(), CategoryID: p.CategoryID().String(), Name: p.Name(), Slug: p.Slug(), SKU: p.SKU(), Description: p.Description(), PriceMinor: p.Price().Amount(), Currency: p.Price().Currency(), ImageURL: p.ImageURL(), Active: p.IsActive()}
}
func InventoryFromDomain(i *inventory.Inventory) InventoryResponse {
	return InventoryResponse{ProductID: i.ProductID().String(), OnHand: i.OnHand(), Reserved: i.Reserved(), Available: i.Available()}
}
func CartFromDomain(c *cart.Cart) CartResponse {
	items := []CartItemResponse{}
	for _, it := range c.Items() {
		sub, _ := it.Subtotal()
		items = append(items, CartItemResponse{ProductID: it.ProductID().String(), ProductName: it.ProductName(), UnitPriceMinor: it.UnitPrice().Amount(), Currency: it.UnitPrice().Currency(), Quantity: it.Quantity(), SubtotalMinor: sub.Amount()})
	}
	total, _ := c.Total()
	return CartResponse{ID: c.ID().String(), Currency: c.Currency(), Items: items, TotalMinor: total.Amount(), TotalQuantity: c.TotalQuantity()}
}
func OrderFromDomain(o *order.Order) OrderResponse {
	items := []OrderItemResponse{}
	for _, it := range o.Items() {
		items = append(items, OrderItemResponse{ProductID: it.ProductID().String(), ProductName: it.ProductName(), UnitPriceMinor: it.UnitPrice().Amount(), Currency: it.UnitPrice().Currency(), Quantity: it.Quantity()})
	}
	return OrderResponse{
		ID:               o.ID().String(),
		UserID:           o.UserID().String(),
		Status:           string(o.Status()),
		Items:            items,
		SubtotalMinor:    o.Subtotal().Amount(),
		ShippingFeeMinor: o.ShippingFee().Amount(),
		DiscountMinor:    o.Discount().Amount(),
		TotalMinor:       o.Total().Amount(),
		Currency:         o.Total().Currency(),
		ShippingAddress:  o.ShippingAddress(),
	}
}
func PaymentFromDomain(p *payment.Transaction) PaymentResponse {
	return PaymentResponse{ID: p.ID().String(), OrderID: p.OrderID().String(), Provider: string(p.Provider()), Status: string(p.Status()), AmountMinor: p.Amount().Amount(), Currency: p.Amount().Currency(), ProviderReference: p.ProviderReference()}
}
