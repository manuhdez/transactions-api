package event

import (
	"context"
)

type Type string

const (
	AccountCreatedType  Type = "event.accounts.account_created"
	WithdrawCreatedType Type = "event.transactions.withdraw_created"
	DepositCreatedType  Type = "event.transactions.deposit_created"
)

type Event interface {
	Type() Type
	Body() []byte
}

type Handler interface {
	Handle(ctx context.Context, event Event) error
}
