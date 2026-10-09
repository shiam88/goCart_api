package response

import (
	"encoding/json"
	"errors"
	authapp "goCartAPI/application/auth"
	orderapp "goCartAPI/application/order"
	"goCartAPI/domain/cart"
	"goCartAPI/domain/category"
	"goCartAPI/domain/inventory"
	"goCartAPI/domain/order"
	"goCartAPI/domain/payment"
	"goCartAPI/domain/product"
	"goCartAPI/domain/user"
	"net/http"
)

type errorBody struct {
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func Error(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	var body errorBody
	body.Error.Code = code
	body.Error.Message = message
	_ = json.NewEncoder(w).Encode(body)
}
func FromError(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	code := "BAD_REQUEST"
	switch {
	case errors.Is(err, authapp.ErrInvalidCredentials):
		status = http.StatusUnauthorized
		code = "INVALID_CREDENTIALS"
	case errors.Is(err, orderapp.ErrOrderForbidden):
		status = http.StatusForbidden
		code = "FORBIDDEN"
	case errors.Is(err, user.ErrUserNotFound), errors.Is(err, product.ErrProductNotFound), errors.Is(err, category.ErrCategoryNotFound), errors.Is(err, cart.ErrCartNotFound), errors.Is(err, order.ErrOrderNotFound), errors.Is(err, payment.ErrTransactionNotFound), errors.Is(err, inventory.ErrInventoryNotFound):
		status = http.StatusNotFound
		code = "NOT_FOUND"
	case errors.Is(err, user.ErrEmailAlreadyUsed), errors.Is(err, product.ErrSKUAlreadyExists):
		status = http.StatusConflict
		code = "CONFLICT"
	case errors.Is(err, inventory.ErrInsufficientStock):
		status = http.StatusConflict
		code = "INSUFFICIENT_STOCK"
	}
	Error(w, status, code, err.Error())

}
