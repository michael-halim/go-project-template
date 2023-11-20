package usecase

import (
	"github.com/michael-halim/go-transactions/internal/transactions"
)

type transactionsScoreHttpUsecase struct {
	usecase transactions.ITransactionsRepository
}

func NewTransactionsScoreHttpUsecase(usecase transactions.ITransactionsRepository) transactions.ITransactionsUsecase {
	return &transactionsScoreHttpUsecase{
		usecase: usecase,
	}
}

func (t *transactionsScoreHttpUsecase) List() error {
	return nil
}
