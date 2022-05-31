package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
)

type Controllers struct {
	Status          controllers.StatusController
	CreateAccount   controllers.CreateAccountController
	FindAccount     controllers.FindAccountController
	FindAllAccounts controllers.FindAllAccountsController
}

var InitControllers = wire.NewSet(
	controllers.NewStatusController,
	controllers.NewCreateAccountController,
	controllers.NewFindAccountController,
	controllers.NewFindAllAccountsControllers,
)

func InitializeControllers(s Services) Controllers {
	return Controllers{
		Status:          controllers.NewStatusController(),
		CreateAccount:   controllers.NewCreateAccountController(s.CreateAccount),
		FindAccount:     controllers.NewFindAccountController(s.FindAccount),
		FindAllAccounts: controllers.NewFindAllAccountsControllers(s.FindAll),
	}
}
