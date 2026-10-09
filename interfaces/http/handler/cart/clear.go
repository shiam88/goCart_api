package cart

import (
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Clear(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.CurrentUser(r.Context())
	c, err := h.service.Clear(r.Context(), u.ID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.CartFromDomain(c))
}
