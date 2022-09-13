package handler

import (
	"context"
	"encoding/json"
	"log"

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
	var data depositCreatedEventBody
	err := json.Unmarshal(e.Body(), &data)
	if err != nil {
		log.Printf("Error parsing created deposit event: %e", err)
	}

	return h.service.Increase(data.Account, data.Amount)
}
