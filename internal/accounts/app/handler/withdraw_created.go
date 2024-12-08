package handler

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
)

var WithdrawCreatedType event.Type = "event.transactions.withdraw_created"

type WithdrawCreated struct {
	service service.DecreaseBalance
}

func NewWithdrawCreated(srv service.DecreaseBalance) WithdrawCreated {
	return WithdrawCreated{srv}
}

func (handler WithdrawCreated) Handle(_ context.Context, e event.Event) error {
	data, err := event.NewWithdrawCreatedBody(e.Body())
	if err != nil {
		fmt.Printf("There was an error parsing the event body: %e", err)
		return err
	}

	return handler.service.Decrease(data.Account, data.Amount)
}
