package transaction_resources

import "time"

type TransactionCreateResponse struct {
	Transaction TransactionCreateResponseTransaction `json:"transaction"`

}

type TransactionCreateResponseTransaction struct {
	Id         int64     `json:"id"`
	ProductId  int64     `json:"product_id"`
	UserId     int64     `json:"user_id"`
	Quantity   int64     `json:"quantity"`
	TotalPrice int64     `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}