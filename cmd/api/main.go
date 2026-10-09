package main

//go clean -modcache -cache
import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	adminapp "goCartAPI/application/admin"
	authapp "goCartAPI/application/auth"
	cartapp "goCartAPI/application/cart"
	categoryapp "goCartAPI/application/category"
	inventoryapp "goCartAPI/application/inventory"
	orderapp "goCartAPI/application/order"
	paymentapp "goCartAPI/application/payment"
	productapp "goCartAPI/application/product"
	"goCartAPI/config"
	"goCartAPI/domain/shared"
	"goCartAPI/domain/user"
	authinfra "goCartAPI/infrastructure/auth"
	"goCartAPI/infrastructure/logging"
	paymentinfra "goCartAPI/infrastructure/payment"
	"goCartAPI/infrastructure/postgres"
	"goCartAPI/infrastructure/postgres/repository"
	httpapi "goCartAPI/interfaces/http"
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

func main() {
	logger := logging.New()

	cfg, err := config.Load()
	if err != nil {
		logger.Error("config", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()
	db, err := postgres.Open(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Error("database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	// Infrastructure implementations.
	userRepo := repository.NewUserRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db)
	inventoryRepo := repository.NewInventoryRepository(db)
	cartRepo := repository.NewCartRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	paymentRepo := repository.NewPaymentRepository(db)
	txManager := postgres.NewTransactionManager(db)

	passwordHasher := authinfra.NewBcryptHasher()

	tokenService := authinfra.NewJWTService(
		cfg.JWTSecret,
		cfg.AccessTokenTTL,
		cfg.RefreshTokenTTL,
	)

	mockGateway := paymentinfra.NewMockGateway()
	codGateway := paymentinfra.NewCODGateway()
	gatewayResolver := paymentinfra.NewProviderFactory(mockGateway, codGateway)

	bootstrapAdmin(ctx, logger, cfg, userRepo, passwordHasher)

	// Application services. Constructors return interfaces, not concrete service structs.
	authService := authapp.NewService(userRepo, passwordHasher, tokenService)
	categoryService := categoryapp.NewService(categoryRepo)
	productService := productapp.NewService(productRepo, categoryRepo, inventoryRepo, txManager)
	inventoryService := inventoryapp.NewService(inventoryRepo)
	cartService := cartapp.NewService(cartRepo, productRepo, cfg.Currency)

	shippingFee, err := shared.NewMoney(cfg.ShippingFee, cfg.Currency)
	if err != nil {
		logger.Error("shipping fee", "error", err)
		os.Exit(1)
	}

	orderService := orderapp.NewService(
		productRepo,
		inventoryRepo,
		cartRepo,
		orderRepo,
		paymentRepo,
		txManager,
		gatewayResolver,
		shippingFee,
	)

	paymentService := paymentapp.NewService(paymentRepo, orderRepo, inventoryRepo, gatewayResolver, txManager)
	adminService := adminapp.NewService(orderRepo, orderService, inventoryService)

	// Middleware dependencies.
	authMiddleware := middleware.NewAuthMiddleware(tokenService)
	middlewares := middleware.NewMiddlewares(authMiddleware)
	manager := middleware.NewManager()
	manager.Use(
		middleware.RequestID,
		middleware.CORS,
		func(next http.Handler) http.Handler {
			return middleware.Logging(logger, next)
		},
		func(next http.Handler) http.Handler {
			return middleware.Recovery(logger, next)
		},
	)

	// Handlers depend on application interfaces (Dependency Inversion).
	handlers := httpapi.Handlers{
		Auth:     auth_handler.NewHandler(authService, middlewares),
		Product:  product_handler.NewHandler(productService, middlewares),
		Category: category_handler.NewHandler(categoryService, middlewares),
		Cart:     cart_handler.NewHandler(cartService, middlewares),
		Order:    order_handler.NewHandler(orderService, middlewares),
		Payment:  payment_handler.NewHandler(paymentService, middlewares),
		Admin:    admin_handler.NewHandler(adminService, middlewares),
		Health:   health_handler.NewHandler(db.Raw()),
	}

	router := httpapi.NewRouter(handlers, manager)
	server := httpapi.NewServer(":"+cfg.HTTPPort, router)

	go func() {
		logger.Info("server_started", "port", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server_failed", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = httpapi.Shutdown(shutdownCtx, server)
	logger.Info("server_stopped")
}

func bootstrapAdmin(
	ctx context.Context,
	logger *slog.Logger,
	cfg config.Config,
	users user.Repository,
	hasher authapp.PasswordHasher,
) {
	if cfg.BootstrapAdminEmail == "" || cfg.BootstrapAdminPassword == "" {
		return
	}

	exists, err := users.EmailExists(ctx, cfg.BootstrapAdminEmail)
	if err != nil || exists {
		return
	}

	hash, err := hasher.Hash(cfg.BootstrapAdminPassword)
	if err != nil {
		return
	}

	u, err := user.NewUser(cfg.BootstrapAdminName, cfg.BootstrapAdminEmail, hash)
	if err != nil {
		return
	}
	_ = u.ChangeRole(user.RoleAdmin)

	if err := users.Save(ctx, u); err != nil {
		logger.Warn("bootstrap_admin_save_failed", "error", err)
	}
}

/*
service:

db:

image:postgres:16-alpine
environment:
POSTGRES_DB:
POSTGRES_USER :
POSTGRES_PASSWORD  :

volumes:
   - goCArtAPI_db:/var/lib/postgresql/data
   - ./internal/infrastructure/postgres/migrations:/docker_entrypoint-initdb.d:ro
   healthcheck:
   test: ["CMD-SHELL", "pg_isready -U goCartAPI"]
   interval :5s
   timeout : 5s
   retrires: 10
   ports:
   - "5432:5432"


   api:
    build: .
	depends_on:
	 condition: service_healthy
	  environment:
*/
