package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type FindAllTransactions struct {
	repository transaction.Repository
}

func NewFindAllTransactionsService(r transaction.Repository) FindAllTransactions {
	return FindAllTransactions{r}
}

func (s FindAllTransactions) Invoke(ctx context.Context) ([]transaction.Transaction, error) {
	return s.repository.FindAll(ctx)
}
