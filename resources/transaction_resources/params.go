package transaction_resources

type TransactionCreateRequest struct {
	ProductId int64 `json:"product_id" validate:"required"`
	Quantity  int64 `json:"quantity" validate:"required,gte=1"`
}
