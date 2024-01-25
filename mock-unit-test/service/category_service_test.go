package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"src/mock-unit-test/entity"
	"src/mock-unit-test/repository"
	"testing"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	// Program Mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)
	category, errs := categoryService.Get("1")
	assert.Nil(t, category, "Category must empty")
	assert.NotNil(t, errs, "Error must returned")
}

func TestCategoryService_GetSuccess(t *testing.T) {
	// Program Mock
	categoryEntity := entity.Category{
		Id:   "1",
		Name: "Pakaian",
	}

	categoryRepository.Mock.On("FindById", "2").Return(categoryEntity)
	category, errs := categoryService.Get("2")
	assert.NotNil(t, category, "Category must returned")
	assert.Nil(t, errs, "Error must empty")
	assert.Equal(t, categoryEntity.Id, category.Id)
	assert.Equal(t, categoryEntity.Name, category.Name)
}
