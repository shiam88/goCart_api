package product

import (
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	if err := h.service.Delete(r.Context(), id); err != nil {
		response.FromError(w, err)
		return
	}
	response.NoContent(w)
}
