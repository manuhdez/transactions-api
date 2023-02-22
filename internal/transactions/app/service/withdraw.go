package service

import (
	"context"
	"errors"
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
	if withdraw.Type != transaction.Withdrawal {
		return errors.New("invalid transaction type")
	}

	err := w.repository.Withdraw(ctx, withdraw)
	if err != nil {
		log.Printf("There was an error executing the withdraw: %e", err)
		return err
	}

	go func() {
		err = w.bus.Publish(ctx, event.NewWithdrawCreated(withdraw.AccountId, withdraw.Amount))
		if err != nil {
			log.Println("error publishing withdraw created event: ", err)
		}
	}()

	return err
}
