package category

import (
	"encoding/json"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	var req dto.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	c, err := h.service.Update(r.Context(), id, req.Name, req.Active)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.CategoryFromDomain(c))
}
