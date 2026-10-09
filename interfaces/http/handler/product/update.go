package product

import (
	"encoding/json"
	productapp "goCartAPI/application/product"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := shared.ParseID(r.PathValue("id"))
	if err != nil {
		response.FromError(w, err)
		return
	}
	var req dto.UpdateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	in := productapp.UpdateInput{ID: id, Name: req.Name, Description: req.Description, ImageURL: req.ImageURL, Active: req.Active}
	if req.CategoryID != "" {
		cid, e := shared.ParseID(req.CategoryID)
		if e != nil {
			response.FromError(w, e)
			return
		}
		in.CategoryID = &cid
	}
	if req.PriceMinor != nil {
		m, e := shared.NewMoney(*req.PriceMinor, req.Currency)
		if e != nil {
			response.FromError(w, e)
			return
		}
		in.Price = &m
	}
	p, err := h.service.Update(r.Context(), in)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, dto.ProductFromDomain(p))
}
