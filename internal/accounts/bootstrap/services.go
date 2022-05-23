package bootstrap

import (
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
)

type Services struct {
	CreateAccount service.CreateService
	FindAll       service.FindAllService
	FindAccount   service.FindAccountService
}

func bootstrapServices(r Repositories) Services {
	return Services{
		CreateAccount: service.NewCreateService(r.Account),
		FindAll:       service.NewFindAllService(r.Account),
		FindAccount:   service.NewFindAccountService(r.Account),
	}
}
