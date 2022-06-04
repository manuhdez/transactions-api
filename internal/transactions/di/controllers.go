package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
)

var InitControllers = wire.NewSet(controllers.NewStatusController)
