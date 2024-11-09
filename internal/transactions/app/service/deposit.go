package service

import (
	"context"
	"fmt"
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
		return fmt.Errorf("[Deposit:Invoke][transaction:%+v][%w]", t, ErrInvalidTransactionType)
	}

	if err := s.repository.Deposit(ctx, t); err != nil {
		return fmt.Errorf("[Deposit:Invoke]%w", err)
	}

	// TODO: what should we do if the publish fails? Where do we handle it?
	if err := s.bus.Publish(ctx, event.NewDepositCreated(t.AccountId, t.Amount)); err != nil {
		log.Printf("[Deposit:Invoke][transaction: %+v]%s", t, err)
	}

	return nil
}
