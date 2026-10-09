package auth

import (
	authapp "goCartAPI/application/auth"
	"goCartAPI/interfaces/http/middleware"
)

type Handler struct {
	service     authapp.Service
	middlewares *middleware.Middlewares
}

func NewHandler(service authapp.Service, middlewares *middleware.Middlewares) *Handler {
	return &Handler{
		service: service, middlewares: middlewares,
	}
}
