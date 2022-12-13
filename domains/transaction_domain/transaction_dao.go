package transaction_domain

import (
	"golang-final-project4-team2/db"
	"golang-final-project4-team2/resources/category_resources"
	"golang-final-project4-team2/resources/transaction_resources"
	"golang-final-project4-team2/utils/error_formats"
	"golang-final-project4-team2/utils/error_utils"
)

var TransactionDomain transactionDomainRepo = &transactionDomain{}

const (
	queryCreateTransaction = `INSERT INTO transactions (product_id, quantity) VALUES($1, $2)
	RETURNING total_price, quantity, product_title`

	queryGetMyTransactions = `SELECT t.id AS t_id, t.product_id AS t_product_id, t.user_id AS t_user_id, 
    quantity, total_price,p.id AS p_id, p.title, p.price, p.stock, p.category_id AS p_category_id, 
    p.created_at AS p_created_at, p.updated_at AS p_updated_at FROM transactions t 
    LEFT JOIN products p ON t.product_id = p.id ORDER BY t.id`

	queryGetUserTransactions = `SELECT t.id AS t_id, t.product_id AS t_product_id, t.user_id AS t_user_id,
	quantity, total_price, p.id AS p_id, p.title, p.price, p.stock. p.category_id AS p_category_id,
	p.created_at AS p_created_at, p.updated_at AS p_updated_at, u.id AS u_id, u.email, u.fullname,
	u.balance, u.created_at AS u_created_at, u.updated_at AS u_updated_at FROM transactions t 
	LEFT JOIN products p ON t.product_id = p.id LEFT JOIN users u ON t.user_id = u.id ORDER BY t.id`
)

type transactionDomainRepo interface {
	CreateTransaction(*transaction_resources.TransactionCreateRequest, string) (*category_resources.CategoriesGetResponse, error_utils.MessageErr)
	GetMyTransactions() (*[]transaction_resources.TransactionGetMyTransactionResponse, error_utils.MessageErr)
	GetUserTransactions() (*[]transaction_resources.TransactionGetUserTransactionsResponse, error_utils.MessageErr)
}

type transactionDomain struct {
}

func (u *transactionDomain) CreateTransaction(req *transaction_resources.TransactionCreateRequest) (*transaction_resources.TransactionCreateResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	row := dbInstance.QueryRow(queryCreateTransaction, req.ProductId, req.Quantity)
	if row.Err() != nil {
		return nil, error_utils.NewBadRequest(row.Err().Error())
	}

	var transaction transaction_resources.TransactionCreateResponse

	err := row.Scan(&transaction.TotalPrice, &transaction.Quantity, &transaction.ProductTitle)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &transaction, nil
}

func (u *transactionDomain) GetMyTransactions() (*[]transaction_resources.TransactionGetMyTransactionResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetMyTransactions)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	transactions := make([]transaction_resources.TransactionGetMyTransactionResponse, 0)
	transactionProducts := make([]transaction_resources.TransactionProductGetResponse, 0)
	var transaction transaction_resources.TransactionGetMyTransactionResponse
	var lastTransactionId int64

	for rows.Next() {
		var transactionProduct transaction_resources.TransactionProductGetResponse

		//	Scan Data
		err := rows.Scan(&transaction.Id, &transaction.ProductId, &transaction.UserId, &transaction.Quantity, &transaction.TotalPrice, &transactionProduct.Id, &transactionProduct.Title, &transactionProduct.Price, &transactionProduct.Stock, &transactionProduct.CategoryId, &transactionProduct.CreatedAt, &transactionProduct.UpdatedAt)

		//	Check wheter product is already in the list
		if transactionProduct.Id != nil {
			var transactionId *int64
			transactionId = new(int64)
			*transactionId = transaction.Id
			transactionProducts = append(transactionProducts, transactionProduct)
		}

		//	If transaction is not the same as the last one, append the transaction to the list
		if lastTransactionId != transaction.Id {
			transaction.Products = transactionProducts
			transactionProducts = make([]transaction_resources.TransactionProductGetResponse, 0)
			transactions = append(transactions, transaction)
		} else {
			//	If transaction is the same as the last one, append the product to the transaction
			transactions[len(transactions)-1].Products = append(transactions[len(transactions)-1].Products, transactionProduct)
			transactionProducts = make([]transaction_resources.TransactionProductGetResponse, 0)
		}

		lastTransactionId = transaction.Id

		if err != nil {
			return nil, error_formats.ParseError(err)
		}
	}
	return &transactions, nil
}

func (u *transactionDomain) GetUserTransactions() (*[]transaction_resources.TransactionGetUserTransactionsResponse, error_utils.MessageErr) {
	dbInstance := db.GetDB()
	rows, err := dbInstance.Query(queryGetMyTransactions)
	if err != nil {
		return nil, error_utils.NewBadRequest(err.Error())
	}
	transactions := make([]transaction_resources.TransactionGetUserTransactionsResponse, 0)
	transactionProducts := make([]transaction_resources.TransactionProductGetResponse, 0)
	transactionUsers := make([]transaction_resources.TransactionUserGetResponse, 0)
	var transaction transaction_resources.TransactionGetUserTransactionsResponse
	var lastTransactionId int64

	for rows.Next() {
		var transactionProduct transaction_resources.TransactionProductGetResponse
		var transactionUser transaction_resources.TransactionUserGetResponse

		//	Scan Data
		err := rows.Scan(&transaction.Id, &transaction.ProductId, &transaction.UserId, &transaction.Quantity, &transaction.TotalPrice, &transactionProduct.Id, &transactionProduct.Title, &transactionProduct.Price, &transactionProduct.Stock, &transactionProduct.CategoryId, &transactionProduct.CreatedAt, &transactionProduct.UpdatedAt, &transactionUser.Id, &transactionUser.Email, &transactionUser.FullName, &transactionUser.Balance, &transactionUser.CreatedAt, &transactionUser.UpdatedAt)

		//	Check wheter product is already in the list
		if transactionProduct.Id != nil {
			var transactionId *int64
			transactionId = new(int64)
			*transactionId = transaction.Id
			transactionProducts = append(transactionProducts, transactionProduct)
		}

		//	Check wheter user is already in the list
		if transactionUser.Id != nil {
			var transactionId *int64
			transactionId = new(int64)
			*transactionId = transaction.Id
			transactionUsers = append(transactionUsers, transactionUser)
		}

		//	If transaction is not the same as the last one, append the transaction to the list
		if lastTransactionId != transaction.Id {
			transaction.Products = transactionProducts
			transaction.Users = transactionUsers
			transactionProducts = make([]transaction_resources.TransactionProductGetResponse, 0)
			transactionUsers = make([]transaction_resources.TransactionUserGetResponse, 0)
			transactions = append(transactions, transaction)
		} else {
			//	If transaction is the same as the last one, append the product to the transaction
			transactions[len(transactions)-1].Products = append(transactions[len(transactions)-1].Products, transactionProduct)
			transactionProducts = make([]transaction_resources.TransactionProductGetResponse, 0)
			transactionUsers = make([]transaction_resources.TransactionUserGetResponse, 0)
		}

		lastTransactionId = transaction.Id

		if err != nil {
			return nil, error_formats.ParseError(err)
		}
	}

	return &transactions, nil
}
