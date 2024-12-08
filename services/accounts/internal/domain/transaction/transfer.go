package transaction

import (
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type Transfer struct {
	UserId domain.ID
	From   account.Account
	To     account.Account
	Amount float32

	events []event.Event
}

// NewTransfer creates a new Transfer instance
func NewTransfer(user domain.ID, from, to account.Account, amount float32) Transfer {
	return Transfer{UserId: user, From: from, To: to, Amount: amount}
}

// CreateTransfer creates a new Transfer instance and registers the domain events
func CreateTransfer(user domain.ID, from, to account.Account, amount float32) Transfer {
	trx := Transfer{UserId: user, From: from, To: to, Amount: amount}

	trx.events = append(trx.events, CreateWithdrawal(from, user, amount).PullEvents()...)
	trx.events = append(trx.events, CreateDeposit(to, user, amount).PullEvents()...)

	return trx
}

// HasValidOwner checks if the user performing the transfer owns the origin account
func (t Transfer) HasValidOwner() bool {
	return t.From.IsOwner(t.UserId)
}

// IsValidAmount checks if the transfer has a valid amount
func (t Transfer) IsValidAmount() error {
	if t.Amount > MaxTransferAmount {
		return ErrTransferAmountTooBig
	}
	if t.Amount <= MinTransferAmount {
		return ErrTransferAmountTooSmall
	}
	if t.Amount > t.From.Balance() {
		return ErrInsufficientBalance
	}
	return nil
}

// Withdrawal returns a withdrawal transaction based on the transfer data
func (t Transfer) Withdrawal() Transaction {
	return CreateWithdrawal(t.From, t.UserId, t.Amount)
}

// Deposit returns a deposit transaction based on the transfer data
func (t Transfer) Deposit() Transaction {
	return CreateDeposit(t.To, t.UserId, t.Amount)
}

// PullEvents returns the generated events and clears the events list
func (t Transfer) PullEvents() []event.Event {
	events := t.events
	t.events = make([]event.Event, 0)
	return events
}
