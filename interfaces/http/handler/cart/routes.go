package cart

import (
	"goCartAPI/interfaces/http/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /api/v1/cart", manager.With(http.HandlerFunc(h.Get), h.middlewares.Authentication))
	mux.Handle("POST /api/v1/cart/items", manager.With(http.HandlerFunc(h.Add), h.middlewares.Authentication))
	mux.Handle("PATCH /api/v1/cart/items/{productId}", manager.With(http.HandlerFunc(h.Update), h.middlewares.Authentication))
	mux.Handle("DELETE /api/v1/cart/items/{productId}", manager.With(http.HandlerFunc(h.Remove), h.middlewares.Authentication))
	mux.Handle("DELETE /api/v1/cart", manager.With(http.HandlerFunc(h.Clear), h.middlewares.Authentication))
}
