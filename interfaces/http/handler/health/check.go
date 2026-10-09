package health

import (
	"context"
	"goCartAPI/interfaces/http/response"
	"net/http"
	"time"
)

func (h *Handler) Check(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()
	if err := h.db.PingContext(ctx); err != nil {
		response.Error(w, 503, "NOT_READY", "database unavailable")
		return
	}
	response.JSON(w, 200, map[string]string{"status": "ok"})
}
