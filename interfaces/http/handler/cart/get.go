package cart

import (
	"goCartAPI/interfaces/http/dto"
	"net/http"

	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.CurrentUser(r.Context())
	c, err := h.service.Get(r.Context(), u.ID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.CartFromDomain(c))
}
