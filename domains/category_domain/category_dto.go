package category_domain

import (
	"time"
)

type Category struct {
	Id                int64     `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int64     `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
