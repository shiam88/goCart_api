package payment

import (
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Initiate(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("orderId"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	ref, err := h.service.Initiate(r.Context(), id)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, map[string]string{"provider_reference": ref})
}
