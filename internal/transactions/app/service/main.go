package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type Depositer interface {
	EventPuller
	Deposit(context.Context, transaction.Transaction) error
}

type Withdrawer interface {
	EventPuller
	Withdraw(context.Context, transaction.Transaction) error
}

type Transferer interface {
	EventPuller
	Transfer(context.Context, transaction.Transfer) error
}

type EventPuller interface {
	PullEvents() []event.Event
}
