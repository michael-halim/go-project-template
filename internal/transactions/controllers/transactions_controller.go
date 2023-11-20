package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/michael-halim/go-transactions/internal/transactions"
)

type transactionsScoreHttpController struct {
	usecase transactions.ITransactionsUsecase
}

func NewTransactionsScoreHttpController(usecase transactions.ITransactionsUsecase) transactions.ITransactionsController {
	return transactionsScoreHttpController{
		usecase: usecase,
	}
}

func (controller transactionsScoreHttpController) List(c *gin.Context) {
	fmt.Println("List transactions")
	c.IndentedJSON(200, gin.H{
		"message": "List transactions",
	})

}
