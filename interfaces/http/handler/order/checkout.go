package order

import (
	"encoding/json"
	orderapp "goCartAPI/application/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/shared"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/middleware"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	auth, _ := middleware.CurrentUser(r.Context())
	var req dto.CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, 400, "INVALID_JSON", "invalid request body")
		return
	}
	//items :=make ([]orderapp.CheckoutItemInput,0,len(req.Items))
	items := make([]orderapp.CheckoutItemInput, 0, len(req.Items))
	for _, it := range req.Items {
		pid, err := shared.ParseID(it.ProductID)
		if err != nil {
			response.FromError(w, err)
			return
		}
		items = append(items, orderapp.CheckoutItemInput{ProductID: pid, Quantity: it.Quantity})
	}
	res, err := h.service.Checkout(r.Context(), orderapp.CheckoutInput{UserID: auth.ID, Source: req.Source, Items: items, ShippingAddress: req.ShippingAddress, PaymentProvider: payment.Provider(req.PaymentMethod)})
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, 201, dto.CheckoutResponse{Order: dto.OrderFromDomain(res.Order), Payment: dto.PaymentFromDomain(res.Payment), ProviderReference: res.ProviderReference})
}
