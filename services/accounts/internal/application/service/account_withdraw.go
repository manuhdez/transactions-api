package service

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type WithdrawService struct {
	trxRepo   transaction.Repository
	publisher event.Bus
}

func NewWithdrawService(trxRepo transaction.Repository, publisher event.Bus) *WithdrawService {
	return &WithdrawService{
		trxRepo:   trxRepo,
		publisher: publisher,
	}
}

// Withdraw withdraws money from an account
func (srv *WithdrawService) Withdraw(ctx context.Context, trx transaction.Transaction) error {
	log.FromContext(ctx).Debug("TransactionService.Withdraw", "withdraw", trx)

	if trx.Type != transaction.Withdrawal {
		return fmt.Errorf("[TransactionService:Withdraw]%w", ErrInvalidTransactionType)
	}

	// TODO: check if account has balance
	// we are going to need to replace the trx entity ids with an account object

	if err := srv.trxRepo.Withdraw(ctx, trx); err != nil {
		return fmt.Errorf("[TransactionService:Withdraw]%w", err)
	}

	for _, ev := range trx.PullEvents() {
		if err := srv.publisher.Publish(ctx, ev); err != nil {
			log.FromContext(ctx).Error("TransactionService.Withdraw", "action", "publish withdraw event", "error", err)
		}
	}

	return nil
}
