package middleware

import (
	"goCartAPI/interfaces/http/response"
	"log/slog"
	"net/http"
)

func Recovery(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if v := recover(); v != nil {
				logger.Error("panic recovered", "panic", v, "path", r.URL.Path)
				response.Error(w, http.StatusInternalServerError, "INTERNAL_ERROR", "internal server error")

			}

		}()
		next.ServeHTTP(w, r)
	})
}
