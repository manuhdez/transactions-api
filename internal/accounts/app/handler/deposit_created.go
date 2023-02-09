package handler

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
)

var DepositCreatedType event.Type = "event.transactions.deposit_created"

type DepositCreated struct {
	service service.IncreaseBalanceService
}

type depositCreatedEventBody struct {
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewHandlerDepositCreated(s service.IncreaseBalanceService) DepositCreated {
	return DepositCreated{s}
}

func (h DepositCreated) Handle(_ context.Context, e event.Event) error {
	data, err := event.NewDepositCreatedBody(e.Body())
	if err != nil {
		fmt.Printf("error parsing event data")
	}

	return h.service.Increase(data.Account, data.Amount)
}
