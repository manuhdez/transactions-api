package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
)

var InitControllers = wire.NewSet(
	controllers.NewStatusController,
	controller.NewDeposit,
	controllers.NewFindAllController,
	controllers.NewWithdrawController,
)
