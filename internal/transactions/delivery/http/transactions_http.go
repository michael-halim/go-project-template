package http

import (
	"github.com/gin-gonic/gin"
	"github.com/michael-halim/go-transactions/internal/transactions"
)

func NewTransactionsHttp(svc transactions.ITransactionsController, router *gin.RouterGroup) {
	transactions_route := router.Group("/transactions")

	transactions_route.GET("/", svc.List)

}
