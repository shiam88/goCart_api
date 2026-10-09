package cart

import (
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Remove(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.CurrentUser(r.Context())
	pid, err := shared.ParseID(r.PathValue("productId"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	c, err := h.service.RemoveItem(r.Context(), u.ID, pid)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.CartFromDomain(c))
}
