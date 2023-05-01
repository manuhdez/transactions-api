//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/users/container"
)

func Init() container.App {
	wire.Build(
		container.Databases,
		container.Buses,
		container.Repositories,
		container.DomainServices,
		container.Services,
		container.Controllers,
		container.Router,
		container.NewApp,
	)

	return container.App{}
}
