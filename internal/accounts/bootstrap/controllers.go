package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
)

var InitControllers = wire.NewSet(
	controllers.NewStatusController,
	controllers.NewCreateAccountController,
	controllers.NewFindAccountController,
	controllers.NewFindAllAccountsControllers,
	controllers.NewDeleteAccountController,
)
