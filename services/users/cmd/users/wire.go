//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func Init() App {
	wire.Build(
		Databases,
		Buses,
		Repositories,
		DomainServices,
		Services,
		Controllers,
		Router,
		NewApp,
	)

	return App{}
}
