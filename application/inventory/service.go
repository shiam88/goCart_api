package inventory

import "goCartAPI/domain/inventory"

type service struct {
	repo inventory.Repository
}

func NewService(repo inventory.Repository) Service {
	return &service{
		repo: repo,
	}
}
