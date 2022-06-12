package service

import (
	"context"
	"errors"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

// Deposit service to handle account deposits.
// It receives a transaction repository to store the deposit.
type Deposit struct {
	repository transaction.Repository
	bus        event.Bus
}

func NewDepositService(r transaction.Repository, b event.Bus) Deposit {
	return Deposit{r, b}
}

func (s Deposit) Invoke(ctx context.Context, t transaction.Transaction) error {
	if t.Type != transaction.Deposit {
		return errors.New("invalid transaction type")
	}

	err := s.repository.Deposit(ctx, t)
	if err != nil {
		return err
	}

	go func() {
		err := s.bus.Publish(ctx, event.DepositCreated{})
		if err != nil {
			log.Println("error publishing deposit created event:", err)
		}
	}()

	return nil
}
