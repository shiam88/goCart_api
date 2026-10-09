package order

import (
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
	"net/http"
	"strconv"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	auth, _ := middleware.CurrentUser(r.Context())
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	res, err := h.service.List(r.Context(), auth.ID, page, limit)
	if err != nil {
		response.FromError(w, err)
		return
	}

	items := make([]dto.OrderResponse, 0, len(res.Items))
	for _, o := range res.Items {
		items = append(items, dto.OrderFromDomain(o))
	}
	response.JSON(w, 200, map[string]any{"items": items, "total": res.Total, "page": res.Page, "limit": res.Limit})
}
