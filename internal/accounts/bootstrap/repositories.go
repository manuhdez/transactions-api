package bootstrap

import (
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

type Repositories struct {
	Account account.Repository
}

func bootstrapRepositories(d Databases) Repositories {
	return Repositories{
		Account: infra.NewAccountMysqlRepository(d.Mysql),
	}
}
