package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/app/handler"
)

var InitHandlers = wire.NewSet(handler.NewHandlerDepositCreated, handler.NewWithdrawCreated)
