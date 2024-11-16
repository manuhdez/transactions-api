package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
)

type AccountCreated struct {
	eventType event.Type
	service   service.CreateAccount
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
	log.Printf("[AccountCreated:Handle][event:%+v]", e)

	data, err := event.NewAccountCreatedBody(e.Body())
	if err != nil {
		return fmt.Errorf("[AccountCreated:Handle][err: %w]", err)
	}

	return h.service.Create(ctx, account.NewAccount(data.Id, data.UserId))
}
