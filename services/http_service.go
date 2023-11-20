package services

import (
	"github.com/michael-halim/go-transactions/internal/transactions"
	transactions_controllers "github.com/michael-halim/go-transactions/internal/transactions/controllers"
	transactions_repository "github.com/michael-halim/go-transactions/internal/transactions/repository"
	transactions_usecase "github.com/michael-halim/go-transactions/internal/transactions/usecase"
	"github.com/michael-halim/go-transactions/pkg"
)

type HttpService struct {
	Transactions transactions.ITransactionsController
}

func NewHttpService() HttpService {
	conn := pkg.NewPostgresConnection()
	transactionsRepository := transactions_repository.NewTransactionsRepository(conn)
	transactionsUsecase := transactions_usecase.NewTransactionsScoreHttpUsecase(transactionsRepository)
	Transactions := transactions_controllers.NewTransactionsScoreHttpController(transactionsUsecase)

	return HttpService{
		Transactions: Transactions,
	}
}
