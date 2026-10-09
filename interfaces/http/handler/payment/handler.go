package payment

import (
	paymentapp "goCartAPI/application/payment"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     paymentapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service paymentapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}
