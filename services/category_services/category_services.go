package category_services

import (
	"golang-final-project3-team2/domains/category_domain"
	"golang-final-project3-team2/resources/category_resources"
	"golang-final-project3-team2/utils/error_utils"
	"golang-final-project3-team2/utils/helpers"
)

var CategoryService categoryServiceRepo = &categoryService{}

type categoryServiceRepo interface {
	CreateCategory(*category_resources.CategoryCreateRequest, string) (*category_resources.CategoryCreateResponse, error_utils.MessageErr)
	GetCategories() (*[]category_resources.CategoriesGetResponse, error_utils.MessageErr)
	UpdateCategory(*category_resources.CategoryUpdateRequest, string) (*category_resources.CategoryUpdateResponse, error_utils.MessageErr)
	DeleteCategory(string) error_utils.MessageErr
}

type categoryService struct{}

func (u *categoryService) CreateCategory(photoReq *category_resources.CategoryCreateRequest, userId string) (*category_resources.CategoryCreateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(photoReq)

	if err != nil {
		return nil, err
	}

	photo, err := category_domain.CategoryDomain.CreateCategory(photoReq, userId)

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *categoryService) GetCategories() (*[]category_resources.CategoriesGetResponse, error_utils.MessageErr) {
	categories, err := category_domain.CategoryDomain.GetCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (u *categoryService) UpdateCategory(request *category_resources.CategoryUpdateRequest, itemId string) (*category_resources.CategoryUpdateResponse, error_utils.MessageErr) {
	updatedCategory, err := category_domain.CategoryDomain.UpdateCategory(request, itemId)

	if err != nil {
		return nil, err
	}

	return updatedCategory, nil
}

func (u *categoryService) DeleteCategory(id string) error_utils.MessageErr {
	err := category_domain.CategoryDomain.DeleteCategory(id)

	if err != nil {
		return err
	}
	return nil
}
