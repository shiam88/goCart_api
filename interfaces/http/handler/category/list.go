package category

import (
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
	"strconv"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	res, err := h.service.List(r.Context(), page, limit)
	if err != nil {
		response.FromError(w, err)
		return
	}
	items := make([]dto.CategoryResponse, 0, len(res.Items))
	for _, c := range res.Items {
		items = append(items, dto.CategoryFromDomain(c))
	}
	response.JSON(w, 200, map[string]any{"items": items, "total": res.Total, "page": res.Page, "limit": res.Limit})
}
