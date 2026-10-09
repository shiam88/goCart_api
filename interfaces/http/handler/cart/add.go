package cart

import (
	"encoding/json"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	u, _ := middleware.CurrentUser(r.Context())
	var req dto.AddCartItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	pid, err := shared.ParseID(req.ProductID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	c, err := h.service.AddItem(r.Context(), u.ID, pid, req.Quantity)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.CartFromDomain(c))
}
