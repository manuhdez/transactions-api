package service

import (
	"context"
	"fmt"
	"log"

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
	events  []event.Event
}

func NewTransactionService(trxRepo transaction.Repository) *TransactionService {
	return &TransactionService{
		trxRepo: trxRepo,
	}
}

// Deposit deposits money into an account
func (srv *TransactionService) Deposit(ctx context.Context, trx transaction.Transaction) error {
	log.Printf("[TransactionService:Deposit][transaction:%+v]", trx)

	if trx.Type != transaction.Deposit {
		return fmt.Errorf("[TransactionService:Deposit]%w", ErrInvalidTransactionType)
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

	// TODO: check if account has balance

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

	// withdraws money from account if the user is authorized and and account has balance
	if err := srv.Withdraw(ctx, withdraw); err != nil {
		return fmt.Errorf("[TransactionService:Transfer]%w", err)
	}

	// Deposit into destination account
	deposit := transaction.NewDeposit(trx.To, trx.UserId, trx.Amount)
	if err := srv.trxRepo.Deposit(ctx, deposit); err != nil {
		return fmt.Errorf("[TransactionService:Transfer]%w", err)
	}

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
