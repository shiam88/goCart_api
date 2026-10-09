package middleware

import (
	"net/http"

	"goCartAPI/domain/user"
)

type Middlewares struct {
	auth *AuthMiddleware
}

func NewMiddlewares(auth *AuthMiddleware) *Middlewares {
	return &Middlewares{
		auth: auth,
	}
}

func (m *Middlewares) Authentication(next http.Handler) http.Handler {
	return m.auth.Require(next)
}

func (m *Middlewares) Admin(next http.Handler) http.Handler {
	return m.auth.Require(RequireRole(user.RoleAdmin, next))
}
