package server

import (
	"github.com/gin-gonic/gin"
	Transactions "github.com/michael-halim/go-transactions/internal/transactions/delivery/http"
	"github.com/michael-halim/go-transactions/services"
)

func (s *Server) MapHandlers(app *gin.Engine) {
	v1 := app.Group("/v1")

	http_service := services.NewHttpService()

	Transactions.NewTransactionsHttp(http_service.Transactions, v1)

}
