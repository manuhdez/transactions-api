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

func NewHandlerDepositCreated(s service.IncreaseBalanceService) DepositCreated {
	return DepositCreated{s}
}

func (h DepositCreated) Handle(_ context.Context, e event.Event) error {
	fmt.Printf("DepositCreated Handler: %s", e.Type())
	return h.service.Increase("", 22)
}
