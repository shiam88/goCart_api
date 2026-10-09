package payment

import (
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	p, err := h.service.Get(r.Context(), id)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 200, dto.PaymentFromDomain(p))
}
