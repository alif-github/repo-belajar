package repository

import "src/mock-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}

type CategoryRepositoryReal struct{}

func (c CategoryRepositoryReal) FindById(_ string) *entity.Category {
	return &entity.Category{
		Id:   "99",
		Name: "Real Cases",
	}
}
