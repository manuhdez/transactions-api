package service

import (
	"context"
	"fmt"
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

func (w Withdraw) Invoke(ctx context.Context, trx transaction.Transaction) error {
	if trx.Type != transaction.Withdrawal {
		return fmt.Errorf("[Withdraw:Invoke][transaction:%+v][err: %w]", trx, ErrInvalidTransactionType)
	}

	if err := w.repository.Withdraw(ctx, trx); err != nil {
		return fmt.Errorf("[Withdraw:Invoke]%w", err)
	}

	if err := w.bus.Publish(ctx, event.NewWithdrawCreated(trx.AccountId, trx.Amount)); err != nil {
		log.Println("error publishing transaction created event: ", err)
	}

	return nil
}
