package auth

import (
	"encoding/json"
	authapp "goCartAPI/application/auth"
	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
	"net/http"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_JSON", "invalid request body")
		return
	}

	res, err := h.service.Login(r.Context(), authapp.LoginInput{Email: req.Email, Password: req.Password})
	if err != nil {
		response.FromError(w, err)
		return
	}
	response.JSON(w, http.StatusOK, dto.AuthResponse{UserID: res.User.ID().String(), Name: res.User.Name(), Email: res.User.Email(), Role: string(res.User.Role()), AccessToken: res.AccessToken, RefreshToken: res.RefreshToken})
}
