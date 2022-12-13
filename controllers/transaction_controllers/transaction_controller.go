package transaction_controllers

import (
	"github.com/gin-gonic/gin"
	"golang-final-project4-team2/resources/transaction_resources"
	"golang-final-project4-team2/services/transaction_services"
	"golang-final-project4-team2/utils/error_utils"
	"net/http"
)

func CreateTransaction(c *gin.Context) {
	var req transaction_resources.TransactionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		theErr := error_utils.NewUnprocessibleEntityError(err.Error())
		c.JSON(theErr.Status(), theErr)
		return
	}

	transaction, err := transaction_services.TransactionService.CreateTransaction(&req)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func GetMyTransactions(c *gin.Context) {
	transactions, err := transaction_services.TransactionService.GetMyTransactions()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func GetUserTransactions(c *gin.Context) {
	transactions, err := transaction_services.TransactionService.GetUserTransactions()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, transactions)
}
