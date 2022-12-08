package product_resources

type ProductCreateRequest struct {
	Title      string `json:"title" validate:"required"`
	Price      int64  `json:"price" validate:"required,gte=0,lte=50000000"`
	Stock      int64  `json:"stock" validate:"required,gte=5"`
	CategoryId int64  `json:"category_id" validate:"required"`
}

type ProductUpdateRequest struct {
	Title      string `json:"title" validate:"required"`
	Price      int64  `json:"price" validate:"required,gte=0,lte=50000000"`
	Stock      int64  `json:"stock" validate:"required,gte=5"`
	CategoryId int64  `json:"category_id" validate:"required"`
}
