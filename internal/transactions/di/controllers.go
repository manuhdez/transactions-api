package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
)

var InitControllers = wire.NewSet(
	controller.NewDeposit,
	controller.NewWithdraw,
	controller.NewFindAllTransactions,
	controller.NewFindAccountTransactions,
)
