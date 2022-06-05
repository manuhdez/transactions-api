package service

import (
	"context"
	"errors"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

// Deposit service to handle account deposits.
// It receives a transaction repository to store the deposit.
type Deposit struct {
	repository transaction.Repository
}

func NewDepositService(r transaction.Repository) Deposit {
	return Deposit{r}
}

func (s Deposit) Invoke(ctx context.Context, t transaction.Transaction) error {
	if t.Type != transaction.Deposit {
		return errors.New("invalid transaction type")
	}

	return s.repository.Deposit(ctx, t)
}
