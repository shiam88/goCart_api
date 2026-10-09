package category

import "goCartAPI/domain/category"

type service struct{ repo category.Repository }

func NewService(repo category.Repository) Service {
	return &service{
		repo: repo,
	}
}
