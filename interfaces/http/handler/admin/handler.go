package admin

import (
	adminapp "goCartAPI/application/admin"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     adminapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service adminapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}
