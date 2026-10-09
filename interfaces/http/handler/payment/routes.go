package payment

import (
	"goCartAPI/interfaces/http/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /api/v1/payments/{orderId}/initiate", manager.With(http.HandlerFunc(h.Initiate), h.middlewares.Authentication))
	mux.HandleFunc("POST /api/v1/payments/callback", h.Callback)
	mux.Handle("GET /api/v1/payments/{id}", manager.With(http.HandlerFunc(h.Get), h.middlewares.Authentication))
	mux.Handle("POST /api/v1/admin/payments/{id}/refund", manager.With(http.HandlerFunc(h.Refund), h.middlewares.Admin))
}
