package repository

import (
	"github.com/stretchr/testify/mock"
	"src/mock-unit-test/entity"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	args := repository.Mock.Called(id)
	if args.Get(0) == nil {
		return nil
	} else {
		category := args.Get(0).(entity.Category)
		return &category
	}
}
