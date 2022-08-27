package handler

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
)

var DepositCreatedType event.Type = "deposit_created"

type DepositCreated struct{}

func (h DepositCreated) Handle(_ context.Context, e event.Event) error {
	log.Println("deposit created event received: %s", e.Type())
	return nil
}
