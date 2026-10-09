package auth

import (
	"goCartAPI/interfaces/http/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.HandleFunc("POST/api/v1/auth/register", h.Register)
	mux.HandleFunc("POST/api/v1/login", h.Login)
	mux.HandleFunc("POST /api/v1/auth/refresh", h.Refresh)

}
