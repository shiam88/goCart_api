package admin

import (
	"goCartAPI/interfaces/http/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /api/v1/admin/orders", manager.With(http.HandlerFunc(h.ListOrders), h.middlewares.Admin))
	mux.Handle("PATCH /api/v1/admin/orders/{id}/status", manager.With(http.HandlerFunc(h.UpdateOrderStatus), h.middlewares.Admin))
	mux.Handle("PATCH /api/v1/admin/products/{id}/stock", manager.With(http.HandlerFunc(h.AdjustStock), h.middlewares.Admin))
}
