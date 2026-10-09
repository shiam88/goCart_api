package cart

import (
	cartapp "goCartAPI/application/cart"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     cartapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service cartapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}
