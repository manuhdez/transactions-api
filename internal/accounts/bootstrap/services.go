package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type Services struct {
	CreateAccount service.CreateService
	FindAll       service.FindAllService
	FindAccount   service.FindAccountService
}

var InitServices = wire.NewSet(
	service.NewCreateService,
	service.NewFindAllService,
	service.NewFindAccountService,
)

func InitializeServices(repo account.Repository) Services {
	return Services{
		CreateAccount: service.NewCreateService(repo),
		FindAll:       service.NewFindAllService(repo),
		FindAccount:   service.NewFindAccountService(repo),
	}
}
