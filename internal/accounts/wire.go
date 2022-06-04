//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

func InitializeServer() bootstrap.Server {
	wire.Build(
		wire.Bind(new(account.Repository), new(infra.AccountMysqlRepository)),
		bootstrap.InitializeRepositories,
		bootstrap.InitializeServices,
		bootstrap.InitializeControllers,
		bootstrap.InitializeServer,
	)
	return bootstrap.Server{}
}
