package product_services

import (
	"golang-final-project4-team2/domains/product_domain"
	"golang-final-project4-team2/resources/product_resources"
	"golang-final-project4-team2/utils/error_utils"
	"golang-final-project4-team2/utils/helpers"
)

var ProductService productServiceRepo = &productService{}

type productServiceRepo interface {
	CreateProduct(*product_resources.ProductCreateRequest) (*product_resources.ProductCreateResponse, error_utils.MessageErr)
	GetProducts() (*[]product_resources.ProductsGetResponse, error_utils.MessageErr)
	UpdateProduct(*product_resources.ProductUpdateRequest, string) (*product_resources.ProductUpdateResponse, error_utils.MessageErr)
	DeleteProduct(string) error_utils.MessageErr
}

type productService struct{}

func (u *productService) CreateProduct(req *product_resources.ProductCreateRequest) (*product_resources.ProductCreateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(req)

	if err != nil {
		return nil, err
	}

	data, err := product_domain.ProductDomain.CreateProduct(req)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (p *productService) GetProducts() (*[]product_resources.ProductsGetResponse, error_utils.MessageErr) {
	products, err := product_domain.ProductDomain.GetProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (u *productService) UpdateProduct(request *product_resources.ProductUpdateRequest, itemId string) (*product_resources.ProductUpdateResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(request)

	if err != nil {
		return nil, err
	}

	updatedProduct, err := product_domain.ProductDomain.UpdateProduct(request, itemId)

	if err != nil {
		return nil, err
	}

	return updatedProduct, nil
}

func (u *productService) DeleteProduct(id string) error_utils.MessageErr {
	err := product_domain.ProductDomain.DeleteProduct(id)

	if err != nil {
		return err
	}
	return nil
}
