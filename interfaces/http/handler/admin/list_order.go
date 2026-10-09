package admin

import (
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
	"strconv"
)

func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	res, err := h.service.ListOrders(r.Context(), page, limit)
	if err != nil {
		response.FromError(w, err)
		return
	}
	items := make([]dto.OrderResponse, 0, len(res.Items))
	for _, o := range res.Items {
		items = append(items, dto.OrderFromDomain(o))
	}
	response.JSON(w, 200, map[string]any{"items": items, "page": res.Page, "limit": res.Limit})
}
