package transactions

type ITransactionsRepository interface {
	List() error
}
