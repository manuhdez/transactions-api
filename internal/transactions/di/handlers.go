package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
)

var InitHandlers = wire.NewSet(handler.NewAccountCreated)
