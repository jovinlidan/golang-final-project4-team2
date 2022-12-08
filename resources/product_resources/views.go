package product_resources

import "time"

type ProductsGetResponse struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryId int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}
type ProductCreateResponse struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryId int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type ProductUpdateResponse struct {
	Product ProductUpdateResponseProduct `json:"product"`
}
type ProductUpdateResponseProduct struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryId int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
