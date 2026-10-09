package product

import (
	"goCartAPI/interfaces/http/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.HandleFunc("GET /api/v1/products", h.List)
	mux.HandleFunc("GET /api/v1/products/{id}", h.Get)
	mux.Handle("POST /api/v1/products", manager.With(http.HandlerFunc(h.Create), h.middlewares.Admin))
	mux.Handle("PATCH /api/v1/products/{id}", manager.With(http.HandlerFunc(h.Update), h.middlewares.Admin))
	mux.Handle("DELETE /api/v1/products/{id}", manager.With(http.HandlerFunc(h.Delete), h.middlewares.Admin))
}
