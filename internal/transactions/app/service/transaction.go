package service

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

const (
	TransactionMinAmount = 10
	DepositMaxAmount     = 5_000
	WithdrawMaxAmount    = 3_000
	TransferMaxAmount    = 10_000
)

var (
	ErrTransferAmountTooLow  = fmt.Errorf("transaction amount must not be lower than %d", TransactionMinAmount)
	ErrTransferAmountTooHigh = fmt.Errorf("transfer amount must not be greater than %d", TransferMaxAmount)
)

type TransactionService struct {
	trxRepo transaction.Repository
	accRepo account.Repository

	events []event.Event
}

func NewTransactionService(trxRepo transaction.Repository, accRepo account.Repository) *TransactionService {
	return &TransactionService{
		trxRepo: trxRepo,
		accRepo: accRepo,
	}
}

// Deposit deposits money into an account
func (srv *TransactionService) Deposit(ctx context.Context, trx transaction.Transaction) error {
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

	srv.pushEvent(event.NewDepositCreated(trx))
	return nil
}

// Withdraw withdraws money from an account
func (srv *TransactionService) Withdraw(ctx context.Context, trx transaction.Transaction) error {
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

	srv.pushEvent(event.NewWithdrawCreated(trx))

	return nil
}

// Transfer transfers money between two accounts
func (srv *TransactionService) Transfer(ctx context.Context, trx transaction.Transfer) error {
	log.Printf("[TransactionService:Transfer][transfer:%+v]", trx)

	// Check amount is between min and max allowed values
	if trx.Amount < TransactionMinAmount {
		return fmt.Errorf("[TransactionService:Transfer][err:%w]", ErrTransferAmountTooLow)
	}

	if trx.Amount > TransferMaxAmount {
		return fmt.Errorf("[TransactionService:Transfer][err:%w]", ErrTransferAmountTooHigh)
	}

	// Withdraw from the origin account
	withdraw := transaction.NewWithdraw(trx.From, trx.UserId, trx.Amount)

	// Check user has access to origin account
	if err := srv.isAccountAuthorized(ctx, &withdraw); err != nil {
		return fmt.Errorf("[TransactionService:Transfer][err: %w]", err)
	}

	// Withdraw from origin account
	if err := srv.trxRepo.Withdraw(ctx, withdraw); err != nil {
		return fmt.Errorf("[TransactionService:Transfer][err: %w]", err)
	}

	// Deposit into destination account
	deposit := transaction.NewDeposit(trx.To, trx.UserId, trx.Amount)
	if err := srv.trxRepo.Deposit(ctx, deposit); err != nil {
		return fmt.Errorf("[TransactionService:Transfer][err: %w]", err)
	}

	srv.pushEvent(event.NewWithdrawCreated(withdraw))
	srv.pushEvent(event.NewDepositCreated(deposit))

	return nil
}

// PullEvents returns the domain events generated and resets the list
func (srv *TransactionService) PullEvents() []event.Event {
	events := srv.events
	srv.events = make([]event.Event, 0)
	return events
}

// pushEvent appends an event to the list of events
func (srv *TransactionService) pushEvent(ev event.Event) {
	srv.events = append(srv.events, ev)
}

// isAccountAuthorized checks if the account exist and is owned by the user
func (srv *TransactionService) isAccountAuthorized(ctx context.Context, trx *transaction.Transaction) error {
	acc, err := srv.accRepo.FindById(ctx, trx.AccountId)
	if err != nil {
		return fmt.Errorf("[isAccountAuthorized]%w", err)
	}

	if acc.UserId != trx.UserId {
		return fmt.Errorf("[isAccountAuthorized][err: %w]", ErrUnauthorized)
	}

	return nil
}
