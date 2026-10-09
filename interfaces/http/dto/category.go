package dto

type CategoryRequest struct {
	Name   string `json:"name"`
	Active *bool  `json:"active,omitempty"`
}
type CategoryResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Active bool   `json:"active"`
}
