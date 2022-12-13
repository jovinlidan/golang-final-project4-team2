package transaction_services

import (
	"golang-final-project4-team2/domains/transaction_domain"
	"golang-final-project4-team2/resources/category_resources"
	"golang-final-project4-team2/resources/transaction_resources"
	"golang-final-project4-team2/utils/error_utils"
	"golang-final-project4-team2/utils/helpers"
)

var TransactionService transactionServiceRepo = &transactionService{}

type transactionServiceRepo interface {
	CreateTransaction(*transaction_resources.TransactionCreateRequest) (*transaction_resources.TransactionCreateResponse, error_utils.MessageErr)
	GetMyTransactions() (*[]transaction_resources.TransactionGetMyTransactionResponse, error_utils.MessageErr)
	GetUserTransactions() (*[]transaction_resources.TransactionGetUserTransactionsResponse, error_utils.MessageErr)
}

type transactionService struct{}

func (u *transactionService) CreateTransaction(photoReq *transaction_resources.TransactionCreateRequest, userId string) (*category_resources.CategoriesGetResponse, error_utils.MessageErr) {
	err := helpers.ValidateRequest(photoReq)

	if err != nil {
		return nil, err
	}

	photo, err := transaction_domain.TransactionDomain.CreateTransaction(photoReq, userId)

	if err != nil {
		return nil, err
	}

	return photo, nil
}

func (u *transactionService) GetMyTransactions() (*[]transaction_resources.TransactionGetMyTransactionResponse, error_utils.MessageErr) {
	categories, err := transaction_domain.TransactionDomain.GetMyTransactions()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (u *transactionService) GetUserTransactions() (*[]transaction_resources.TransactionGetUserTransactionsResponse, error_utils.MessageErr) {
	categories, err := transaction_domain.TransactionDomain.GetUserTransactions()

	if err != nil {
		return nil, err
	}

	return categories, nil
}
