package bootstrap

import "github.com/manuhdez/transactions-api/internal/accounts/controllers"

type Controllers struct {
	Status          controllers.StatusController
	CreateAccount   controllers.CreateAccountController
	FindAccount     controllers.FindAccountController
	FindAllAccounts controllers.FindAllAccountsController
}

func InitializeControllers(s Services) Controllers {
	return Controllers{
		Status:          controllers.NewStatusController(),
		CreateAccount:   controllers.NewCreateAccountController(s.CreateAccount),
		FindAccount:     controllers.NewFindAccountController(s.FindAccount),
		FindAllAccounts: controllers.NewFindAllAccountsControllers(s.FindAll),
	}
}
