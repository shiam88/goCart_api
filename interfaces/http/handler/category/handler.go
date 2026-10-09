package category

import (
	categoryapp "goCartAPI/application/category"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     categoryapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service categoryapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{service: service, middlewares: middlewares}
}
