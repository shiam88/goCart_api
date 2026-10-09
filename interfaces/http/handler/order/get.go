package order

import (
	"goCartAPI/domain/shared"
	"goCartAPI/domain/user"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	auth, _ := middleware.CurrentUser(r.Context())
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	o, err := h.service.Get(r.Context(), id, auth.ID, auth.Role == user.RoleAdmin)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.OrderFromDomain(o))
}
