package middleware

import (
	"goCartAPI/domain/user"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func RequireRole(role user.Role, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		u, ok := CurrentUser(r.Context())
		if !ok {
			response.Error(w, http.StatusUnauthorized, "UNAUTHORIZED", "authentication required")
			return
		}
		if u.Role != role {
			response.Error(w, http.StatusForbidden, "FORBIDDEN", "insufficient permission")
			return
		}
		next.ServeHTTP(w, r)
	})
}
