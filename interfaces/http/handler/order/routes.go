package order

import (
	"goCartAPI/interfaces/http/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /api/v1/orders/checkout", manager.With(http.HandlerFunc(h.Checkout), h.middlewares.Authentication))
	mux.Handle("GET /api/v1/orders", manager.With(http.HandlerFunc(h.List), h.middlewares.Authentication))
	mux.Handle("GET /api/v1/orders/{id}", manager.With(http.HandlerFunc(h.Get), h.middlewares.Authentication))
	//mux.Handle("POST /api/v1/orders/{id}/cancel", manager.With(http.HandlerFunc(h.Cancel), h.middlewares.Authentication))
}
