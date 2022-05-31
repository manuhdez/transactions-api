package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

type Repositories struct {
	Account account.Repository
}

var InitializeRepositories = wire.NewSet(
	InitializeDB,
	wire.Bind(new(account.Repository), new(infra.AccountMysqlRepository)),
	infra.NewAccountMysqlRepository,
)
