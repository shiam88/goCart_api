package admin

import (
	"encoding/json"
	"goCartAPI/domain/order"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	var req dto.UpdateOrderStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	o, err := h.service.UpdateOrderStatus(r.Context(), id, order.Status(req.Status))
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.OrderFromDomain(o))
}
