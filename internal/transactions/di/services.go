package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

var (
	InitServices = wire.NewSet(service.NewDepositService, service.NewFindAllTransactionsService)
)
