package transaction_resources

import "time"

type TransactionCreateResponse struct {
	TotalPrice   int64  `json:"total_price"`
	Quantity     int64  `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionUserGetResponse struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

type TransactionProductGetResponse struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Price      int64     `json:"price"`
	Stock      int64     `json:"stock"`
	CategoryId int64     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type TransactionGetMyTransactionResponse struct {
	Id         int64                           `json:"id"`
	ProductId  int64                           `json:"product_id"`
	UserID     int64                           `json:"user_id"`
	Quantity   int64                           `json:"quantity"`
	TotalPrice int64                           `json:"total_price"`
	Products   []TransactionProductGetResponse `json:"products"`
}

type TransactionGetUserTransactionsResponse struct {
	Id         int64                           `json:"id"`
	ProductId  int64                           `json:"product_id"`
	UserID     int64                           `json:"user_id"`
	Quantity   int64                           `json:"quantity"`
	TotalPrice int64                           `json:"total_price"`
	Products   []TransactionProductGetResponse `json:"products"`
	Users      []TransactionUserGetResponse    `json:"users"`
}
