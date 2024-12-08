package service

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type Transfer struct {
	publisher event.Bus
	trxRepo   transaction.Repository
}

func NewTransferService(trx transaction.Repository, pub event.Bus) Transfer {
	return Transfer{
		trxRepo:   trx,
		publisher: pub,
	}
}

// Transfer handles money transfers between two accounts
func (srv Transfer) Transfer(ctx context.Context, trx transaction.Transfer) error {
	log.FromContext(ctx).Debug("[TransferService:Transfer]", "userID", trx.UserId, "from", trx.From.Id, "to", trx.To.Id, "amount", trx.Amount)

	// check if user performing the transfer owns the origin account
	if !trx.HasValidOwner() {
		return ErrUnauthorizedTransaction
	}

	//  validate amount and account balance
	if err := trx.IsValidAmount(); err != nil {
		return fmt.Errorf("[TransferService:Transfer][err: %w]", err)
	}

	// persist transactions in storage
	if err := srv.trxRepo.Transfer(ctx, trx); err != nil {
		return fmt.Errorf("[TransferService:Transfer]%w", err)
	}

	// publish domain events
	for _, e := range trx.PullEvents() {
		if err := srv.publisher.Publish(ctx, e); err != nil {
			log.FromContext(ctx).Error("[TransferService:Transfer]", "event", e.Type(), "error", err)
		}
	}

	return nil
}
