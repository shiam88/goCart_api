package admin

import (
	"encoding/json"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

type adjustStockRequest struct {
	Delta int `json:"delta"`
}

func (h *Handler) AdjustStock(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	var req adjustStockRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	inv, err := h.service.AdjustStock(r.Context(), id, req.Delta)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.InventoryFromDomain(inv))
}
