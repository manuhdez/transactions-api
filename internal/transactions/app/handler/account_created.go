package handler

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
)

type AccountCreated struct {
	eventType event.Type
	service   service.CreateAccount
}

type accountCreatedBody struct {
	Id string `json:"body"`
}

func NewAccountCreated(s service.CreateAccount) AccountCreated {
	return AccountCreated{
		eventType: event.AccountCreatedType,
		service:   s,
	}
}

func (h AccountCreated) Type() event.Type {
	return h.eventType
}

func (h AccountCreated) Handle(ctx context.Context, e event.Event) error {
	data, err := event.NewAccountCreatedBody(e.Body())
	if err != nil {
		fmt.Printf("error parsing event body")
	}
	return h.service.Create(ctx, account.NewAccount(data.Id))
}
