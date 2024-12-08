package service

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type DepositService struct {
	trxRepo   transaction.Repository
	publisher event.Bus
}

func NewDepositService(trxRepo transaction.Repository, publisher event.Bus) *DepositService {
	return &DepositService{
		trxRepo:   trxRepo,
		publisher: publisher,
	}
}

// Deposit deposits money into an account
func (srv *DepositService) Deposit(ctx context.Context, trx transaction.Transaction) error {
	log.FromContext(ctx).Debug("DepositService.Deposit", "deposit", trx)

	if trx.Type != transaction.Deposit {
		return fmt.Errorf("[TransactionService:Deposit]%w", ErrInvalidTransactionType)
	}

	// TODO: check if account has balance

	if err := srv.trxRepo.Deposit(ctx, trx); err != nil {
		return fmt.Errorf("[TransactionService:Deposit]%w", err)
	}

	for _, ev := range trx.PullEvents() {
		if err := srv.publisher.Publish(ctx, ev); err != nil {
			log.FromContext(ctx).Error("DepositService.Deposit", "action", "publish deposit event", "error", err)
		}
	}

	return nil
}
