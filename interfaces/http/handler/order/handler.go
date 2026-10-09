package order

import (
	orderapp "goCartAPI/application/order"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     orderapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service orderapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}
