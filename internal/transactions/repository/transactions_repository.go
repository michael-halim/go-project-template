package repository

import (
	"github.com/michael-halim/go-transactions/internal/transactions"
	"gorm.io/gorm"
)

type transactionsRepository struct {
	conn *gorm.DB
}

func NewTransactionsRepository(conn *gorm.DB) transactions.ITransactionsRepository {
	return &transactionsRepository{
		conn: conn,
	}
}

func (t *transactionsRepository) List() error {
	return nil
}
