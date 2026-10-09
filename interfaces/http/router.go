package httpapi

import (
	"net/http"

	admin_handler "goCartAPI/interfaces/http/handler/admin"
	auth_handler "goCartAPI/interfaces/http/handler/auth"
	cart_handler "goCartAPI/interfaces/http/handler/cart"
	category_handler "goCartAPI/interfaces/http/handler/category"
	health_handler "goCartAPI/interfaces/http/handler/health"
	order_handler "goCartAPI/interfaces/http/handler/order"
	payment_handler "goCartAPI/interfaces/http/handler/payment"
	product_handler "goCartAPI/interfaces/http/handler/product"
	"goCartAPI/interfaces/http/middleware"
)

type Handlers struct {
	Auth     *auth_handler.Handler
	Product  *product_handler.Handler
	Category *category_handler.Handler
	Cart     *cart_handler.Handler
	Order    *order_handler.Handler
	Payment  *payment_handler.Handler
	Admin    *admin_handler.Handler
	Health   *health_handler.Handler
}

func NewRouter(h Handlers, manager *middleware.Manager) http.Handler {
	mux := http.NewServeMux()

	h.Health.RegisterRoutes(mux)
	h.Auth.RegisterRoutes(mux, manager)
	h.Product.RegisterRoutes(mux, manager)
	h.Category.RegisterRoutes(mux, manager)
	h.Cart.RegisterRoutes(mux, manager)
	h.Order.RegisterRoutes(mux, manager)
	h.Payment.RegisterRoutes(mux, manager)
	h.Admin.RegisterRoutes(mux, manager)

	return manager.Wrap(mux)
}
