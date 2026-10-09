package product

import "errors"

var (
	ErrInvalidProductID  = errors.New("invalid product id")
	ErrInvalidCategoryID = errors.New("invalid category id")
	ErrInvalidName       = errors.New("product name is required")
	ErrInvalidSKU        = errors.New("product sku is required")
	ErrInvalidPrice      = errors.New("product price must be greater than zero")
	ErrProductNotFound   = errors.New("product not found")
	ErrSKUAlreadyExists  = errors.New("product sku already exists")
)
