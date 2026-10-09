package product

import (
	"encoding/json"
	productapp "goCartAPI/application/product"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"net/http"

	"goCartAPI/interfaces/http/response"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	cid, err := shared.ParseID(req.CategoryID)
	if err != nil {
		response.FromError(w, err)
		return
	}
	money, err := shared.NewMoney(req.PriceMinor, req.Currency)
	if err != nil {
		response.FromError(w, err)
		return
	}
	p, err := h.service.Create(r.Context(), productapp.CreateInput{CategoryID: cid, Name: req.Name, SKU: req.SKU, Description: req.Description, Price: money, ImageURL: req.ImageURL, InitialStock: req.InitialStock})
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, dto.ProductFromDomain(p))
}
