package transaction

import (
	"time"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type Type string

const (
	Deposit    Type = "deposit"
	Withdrawal Type = "withdrawal"
)

const (
	TransactionMinAmount = 10
	DepositMaxAmount     = 5_000
	WithdrawMaxAmount    = 3_000

	MaxTransferAmount = 10_000
	MinTransferAmount = 0
)

type Transaction struct {
	Type      Type
	AccountId string
	UserId    string
	Amount    float32
	Date      time.Time

	events []event.Event
}

// NewTransaction creates a new transaction instance
func NewTransaction(t Type, acc, user string, amount float32) Transaction {
	return Transaction{Type: t, AccountId: acc, UserId: user, Amount: amount}
}

// NewDeposit creates a new deposit transaction instance
func NewDeposit(acc, user string, amount float32) Transaction {
	return Transaction{Type: Deposit, AccountId: acc, UserId: user, Amount: amount}
}

// NewWithdraw creates a new withdraw transaction instance
func NewWithdraw(acc, user string, amount float32) Transaction {
	return Transaction{Type: Withdrawal, AccountId: acc, UserId: user, Amount: amount}
}

// CreateDeposit creates a deposit transaction and registers a new created deposit event
func CreateDeposit(acc account.Account, userID domain.ID, amount float32) Transaction {
	trx := NewDeposit(acc.Id(), userID.String(), amount)
	trx.events = append(trx.events, NewDepositCreated(trx))
	return trx
}

// CreateWithdrawal creates a withdraw transaction and registers a new created withdraw event
func CreateWithdrawal(acc account.Account, userID domain.ID, amount float32) Transaction {
	trx := NewWithdraw(acc.Id(), userID.String(), amount)
	trx.events = append(trx.events, NewWithdrawCreated(trx))
	return trx
}

// PullEvents returns the events associated with the transaction
func (t Transaction) PullEvents() []event.Event {
	events := t.events
	t.events = []event.Event{}
	return events
}
