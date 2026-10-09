package product

import (
	productapp "goCartAPI/application/product"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     productapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service productapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}
