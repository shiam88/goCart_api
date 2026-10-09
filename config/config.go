package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort               string
	DatabaseURL            string
	JWTSecret              string
	AccessTokenTTL         time.Duration
	RefreshTokenTTL        time.Duration
	Currency               string
	ShippingFee            int64
	BootstrapAdminEmail    string
	BootstrapAdminPassword string
	BootstrapAdminName     string
}

func Load() (Config, error) {
	_ = godotenv.Load()
	cfg := Config{
		HTTPPort:               getenv("HTTP_PORT", "3000"),
		DatabaseURL:            getenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/commercecore?sslmode=disable"),
		JWTSecret:              os.Getenv("JWT_SECRET_KEY"),
		Currency:               getenv("CURRENCY", "BDT"),
		BootstrapAdminEmail:    os.Getenv("BOOTSTRAP_ADMIN_EMAIL"),
		BootstrapAdminPassword: os.Getenv("BOOTSTRAP_ADMIN_PASSWORD"),
		BootstrapAdminName:     getenv("BOOTSTRAP_ADMIN_NAME", "CommerceCore Admin"),
	}
	if cfg.JWTSecret == "" {
		return Config{}, fmt.Errorf("JWT_SECRET_KEY is required")
	}
	accessMins, err := strconv.Atoi(getenv("ACCESS_TOKEN_MINUTES", "60"))
	if err != nil {
		return Config{}, fmt.Errorf("ACCESS_TOKEN_MINUTES: %w", err)
	}
	refreshHours, err := strconv.Atoi(getenv("REFRESH_TOKEN_HOURS", "168"))
	if err != nil {
		return Config{}, fmt.Errorf("REFRESH_TOKEN_HOURS: %w", err)
	}
	cfg.AccessTokenTTL = time.Duration(accessMins) * time.Minute
	cfg.RefreshTokenTTL = time.Duration(refreshHours) * time.Hour
	fee, err := strconv.ParseInt(getenv("SHIPPING_FEE_MINOR", "6000"), 10, 64)
	if err != nil {
		return Config{}, fmt.Errorf("SHIPPING_FEE_MINOR: %w", err)
	}
	cfg.ShippingFee = fee
	return cfg, nil
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
