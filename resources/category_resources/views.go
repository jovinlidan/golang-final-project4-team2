package category_resources

import "time"

type CategoryCreateResponse struct {
	Id        int64     `json:"id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type CategoriesTaskGetResponse struct {
	Id          *int64     `json:"id"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	UserId      *int64     `json:"user_id"`
	CategoryId  *int64     `json:"category_id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type CategoriesGetResponse struct {
	Id        int64                       `json:"id"`
	Type      string                      `json:"type"`
	CreatedAt time.Time                   `json:"created_at"`
	UpdatedAt time.Time                   `json:"updated_at"`
	Tasks     []CategoriesTaskGetResponse `json:"tasks"`
}

type CategoryUpdateResponse struct {
	Id                int64     `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int64     `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}
