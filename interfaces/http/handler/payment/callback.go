package payment

import (
	"encoding/json"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Callback(w http.ResponseWriter, r *http.Request) {
	var req dto.PaymentCallbackRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	id, err := shared.ParseID(req.OrderID)
	if err != nil {
		response.FromError(w, err)
		return
	}

	if err := h.service.Confirm(r.Context(), id, req.ProviderReference, req.Success); err != nil {
		response.FromError(w, err)
		return
	}

	response.JSON(w, 200, map[string]bool{"ok": true})
}
