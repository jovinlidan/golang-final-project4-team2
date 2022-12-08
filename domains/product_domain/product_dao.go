package product_domain

import (
	"golang-final-project4-team2/db"
	"golang-final-project4-team2/resources/product_resources"
	"golang-final-project4-team2/utils/error_formats"
	"golang-final-project4-team2/utils/error_utils"
)

var ProductDomain productDomainRepo = &productDomain{}

const (
	queryCreateProduct = `insert into products(title, price, stock, category_id) values ($1, $2, $3, $4) 
							RETURNING id, title, price, stock, category_id,created_at`

	queryGetProducts = `
	select id, title, price, stock, category_id, created_at from products`

	queryProductUpdate = `UPDATE products set updated_at = now(), title = $1, price = $2, stock = $3, category_id = $4 where id = $5 
                           RETURNING id, title, price, stock, category_id, created_at, updated_at`

	queryDeleteTransactionHistoriesByProductId = `DELETE FROM transaction_histories where product_id = $1`
	queryDeleteProduct                         = `DELETE from products where id = $1`
)

type productDomainRepo interface {
	CreateProduct(*product_resources.ProductCreateRequest) (*product_resources.ProductCreateResponse, error_utils.MessageErr)
	GetProducts() (*[]product_resources.ProductsGetResponse, error_utils.MessageErr)
	UpdateProduct(*product_resources.ProductUpdateRequest, string) (*product_resources.ProductUpdateResponse, error_utils.MessageErr)
	DeleteProduct(string) error_utils.MessageErr
}

type productDomain struct {
}

func (u *productDomain) CreateProduct(req *product_resources.ProductCreateRequest) (*product_resources.ProductCreateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateProduct, req.Title, req.Price, req.Stock, req.CategoryId)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var product product_resources.ProductCreateResponse

	err := row.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &product.CreatedAt)

	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &product, nil
}

func (u *productDomain) GetProducts() (*[]product_resources.ProductsGetResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetProducts)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	products := make([]product_resources.ProductsGetResponse, 0)
	for rows.Next() {
		var product product_resources.ProductsGetResponse

		// Scan Data
		err = rows.Scan(&product.Id, &product.Title, &product.Price, &product.Stock, &product.CategoryId, &product.CreatedAt)
		products = append(products, product)
		if err != nil {
			return nil, error_formats.ParseError(err)
		}
	}

	return &products, nil
}

func (u *productDomain) UpdateProduct(request *product_resources.ProductUpdateRequest, id string) (*product_resources.ProductUpdateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryProductUpdate, request.Title, request.Price, request.Stock, request.CategoryId, id)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var product product_resources.ProductUpdateResponse
	var productProduct product_resources.ProductUpdateResponseProduct

	err := row.Scan(&productProduct.Id, &productProduct.Title, &productProduct.Price, &productProduct.Stock, &productProduct.CategoryId, &productProduct.CreatedAt, &productProduct.UpdatedAt)
	product.Product = productProduct

	if err != nil {
		return nil, error_formats.ParseError(err)
	}
	return &product, nil
}

func (u *productDomain) DeleteProduct(id string) (error error_utils.MessageErr) {
	dbInstance, err := db.GetDB().Begin()
	if err != nil {
		error = error_utils.NewInternalServerError(err.Error())
		return
	}

	defer func() {
		if error != nil {
			dbInstance.Rollback()
			return
		}
		err := dbInstance.Commit()
		if err != nil {
			error = error_utils.NewInternalServerError(err.Error())
		}
	}()

	rows, err := dbInstance.Query(queryDeleteTransactionHistoriesByProductId, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	rows, err = dbInstance.Query(queryDeleteProduct, id)
	if rows.Err() != nil {
		return error_utils.NewBadRequest(err.Error())
	}
	rows.Close()

	return nil
}
