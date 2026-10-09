package auth

import (
	"net/http"

	"goCartAPI/interfaces/http/response"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	if err := h.service.Logout(r.Context()); err != nil {
		response.FromError(w, err)
		return
	}
	response.NoContent(w)
}
