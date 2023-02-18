package service

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type Withdraw struct {
	repository transaction.Repository
	bus        event.Bus
}

func NewWithdrawService(r transaction.Repository, b event.Bus) Withdraw {
	return Withdraw{r, b}
}

func (w Withdraw) Invoke(ctx context.Context, withdraw transaction.Transaction) error {
    err := w.repository.Withdraw(ctx, withdraw)
    if err != nil {
        log.Printf("There was an error executing the withdraw: %e", err)
        return err
    }
	return nil
}
