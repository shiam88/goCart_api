package product

import (
	productapp "goCartAPI/application/product"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
	"strconv"
)

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	limit, _ := strconv.Atoi(q.Get("limit"))
	var cid *shared.ID
	if v := q.Get("category_id"); v != "" {
		id, err := shared.ParseID(v)
		if err != nil {
			response.FromError(w, err)
			return
		}
		cid = &id
	}
	var min, max *int64
	if v := q.Get("min_price"); v != "" {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			response.Error(w, 400, "INVALID_MIN_PRICE", "min_price must be an integer")
			return
		}
		min = &n
	}
	if v := q.Get("max_price"); v != "" {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			response.Error(w, 400, "INVALID_MAX_PRICE", "max_price must be an integer")
			return
		}
		max = &n
	}
	res, err := h.service.List(r.Context(), productapp.ListInput{Page: page, Limit: limit, Search: q.Get("search"), CategoryID: cid, MinPrice: min, MaxPrice: max, OnlyActive: true})
	if err != nil {
		response.FromError(w, err)
		return
	}
	items := make([]dto.ProductResponse, 0, len(res.Items))
	for _, p := range res.Items {
		items = append(items, dto.ProductFromDomain(p))
	}
	response.JSON(w, http.StatusOK, map[string]any{"items": items, "total": res.Total, "page": res.Page, "limit": res.Limit})
}
