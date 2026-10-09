package middleware

import (
	"context"
	"net/http"
	"strings"

	authapp "goCartAPI/application/auth"
	"goCartAPI/domain/shared"
	"goCartAPI/domain/user"
	"goCartAPI/interfaces/http/response"
)

type authKey struct {
}

type AuthUser struct {
	ID   shared.ID
	Role user.Role
}

type AuthMiddleware struct {
	tokens authapp.TokenService
}

func NewAuthMiddleware(tokens authapp.TokenService) *AuthMiddleware {
	return &AuthMiddleware{tokens: tokens}
}
func (m *AuthMiddleware) Require(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := strings.TrimSpace(r.Header.Get("Authorization"))
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			response.Error(w, 401, "UNAUTHORIZED", "missing or invalid Authorization header")
			return
		}

		claims, err := m.tokens.ParseAccessToken(strings.TrimSpace(parts[1]))
		if err != nil {
			response.Error(w, 401, "UNAUTHORIZED", "invalid or expired token")
			return
		}

		ctx := context.WithValue(r.Context(), authKey{}, AuthUser{ID: claims.UserID,
			Role: claims.Role,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
func CurrentUser(ctx context.Context) (AuthUser, bool) {
	v, ok := ctx.Value(authKey{}).(AuthUser)
	return v, ok
}
