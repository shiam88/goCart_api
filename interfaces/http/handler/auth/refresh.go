package auth

import (
	"encoding/json"
	"net/http"

	"goCartAPI/interfaces/http/dto"
	"goCartAPI/interfaces/http/response"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "INVALID_JSON", "invalid request body")
		return
	}
	res, err := h.service.Refresh(r.Context(), req.RefreshToken)
	if err != nil {
		response.Error(w, http.StatusUnauthorized, "INVALID_REFRESH_TOKEN", err.Error())
		return
	}
	response.JSON(w, http.StatusOK, dto.RefreshResponse{AccessToken: res.AccessToken})
}
