package auth

import (
	"encoding/json"
	"net/http"

	authapp "goCartAPI/application/auth"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_JSON", "invalid request body")
		return
	}
	u, err := h.service.Register(r.Context(), authapp.RegisterInput{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	},
	)
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusCreated, u)

}
