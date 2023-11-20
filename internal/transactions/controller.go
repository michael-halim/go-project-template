package transactions

import "github.com/gin-gonic/gin"

type ITransactionsController interface {
	List(c *gin.Context)
}
