package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/controller"
)

var InitControllers = wire.NewSet(
	controller.NewCreateAccount,
	controller.NewFindAccountController,
	controller.NewFindAllAccounts,
	controller.NewDeleteAccount,
)
