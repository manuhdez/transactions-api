package service

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type TransactionService struct {
	trxRepo  transaction.Repository
	accRepo  account.Repository
	eventBus event.Bus
}

func NewTransactionService(trxRepo transaction.Repository, accRepo account.Repository, bus event.Bus) TransactionService {
	return TransactionService{
		trxRepo:  trxRepo,
		accRepo:  accRepo,
		eventBus: bus,
	}
}

// Deposit deposits money into an account
func (srv TransactionService) Deposit(ctx context.Context, trx transaction.Transaction) error {
	log.Printf("[TransactionService:Deposit][transaction:%+v]", trx)

	if trx.Type != transaction.Deposit {
		return fmt.Errorf("[TransactionService:Deposit]%w", ErrInvalidTransactionType)
	}

	if err := srv.isAccountAuthorized(ctx, &trx); err != nil {
		return fmt.Errorf("[TransactionService:Deposit]%w", err)
	}

	if err := srv.trxRepo.Deposit(ctx, trx); err != nil {
		return fmt.Errorf("[TransactionService:Deposit]%w", err)
	}

	if err := srv.eventBus.Publish(ctx, event.NewDepositCreated(trx)); err != nil {
		return fmt.Errorf("[TransactionService:Deposit]%w", err)
	}

	return nil
}

// Withdraw withdraws money from an account
func (srv TransactionService) Withdraw(ctx context.Context, trx transaction.Transaction) error {
	log.Printf("[TransactionService:Withdraw][transaction:%+v]", trx)

	if trx.Type != transaction.Withdrawal {
		return fmt.Errorf("[TransactionService:Withdraw]%w", ErrInvalidTransactionType)
	}

	if err := srv.isAccountAuthorized(ctx, &trx); err != nil {
		return fmt.Errorf("[TransactionService:Withdraw]%w", err)
	}

	if err := srv.trxRepo.Withdraw(ctx, trx); err != nil {
		return fmt.Errorf("[TransactionService:Withdraw]%w", err)
	}

	if err := srv.eventBus.Publish(ctx, event.NewWithdrawCreated(trx.AccountId, trx.Amount)); err != nil {
		log.Println("error publishing transaction created event: ", err)
		return fmt.Errorf("[TransactionService:Withdraw]%w", err)
	}

	return nil
}

// Transfer transfers money between two accounts
func (srv TransactionService) Transfer(ctx context.Context, trx transaction.Transfer) error {
	return fmt.Errorf("[TransactionService:Transfer][err: method not implemented]")
}

// isAccountAuthorized checks if the account exist and is owned by the user and attaches it to the transaction
func (srv TransactionService) isAccountAuthorized(ctx context.Context, trx *transaction.Transaction) error {
	acc, err := srv.accRepo.FindById(ctx, trx.AccountId)
	if err != nil {
		return fmt.Errorf("[isAccountAuthorized]%w", err)
	}

	if acc.UserId != trx.UserId {
		return fmt.Errorf("[isAccountAuthorized][err: %w]", ErrUnauthorized)
	}

	return nil
}
