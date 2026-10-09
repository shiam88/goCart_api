package category

import (
	"encoding/json"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CategoryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	c, err := h.service.Create(r.Context(), req.Name)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 201, dto.CategoryFromDomain(c))
}
