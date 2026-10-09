package category

import "errors"

var (
	ErrInvalidCategoryID = errors.New("invalid category id")
	ErrInvalidName       = errors.New("category name is required")
	ErrCategoryNotFound  = errors.New("category not found")
)
