package category_resources

import "time"

type CategoryCreateResponse struct {
	Id                int64     `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int64     `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type CategoriesProductGetResponse struct {
	Id        *int64     `json:"id"`
	Title     *string    `json:"title"`
	Price     *int64     `json:"price"`
	Stock     *int64     `json:"stock"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type CategoriesGetResponse struct {
	Id                int64                          `json:"id"`
	Type              string                         `json:"type"`
	SoldProductAmount int64                          `json:"sold_product_amount"`
	CreatedAt         time.Time                      `json:"created_at"`
	UpdatedAt         time.Time                      `json:"updated_at"`
	Products          []CategoriesProductGetResponse `json:"products"`
}

type CategoryUpdateResponse struct {
	Id                int64     `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int64     `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}
