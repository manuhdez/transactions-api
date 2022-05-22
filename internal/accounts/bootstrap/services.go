package bootstrap

import (
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
)

type Services struct {
	CreateAccount service.CreateService
}

func bootstrapServices(r Repositories) Services {
	return Services{
		CreateAccount: service.NewCreateService(r.Account),
	}

}
