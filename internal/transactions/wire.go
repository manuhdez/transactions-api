//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/di"
)

func InitServer() *gin.Engine {
	wire.Build(
		di.NewDBConnection,
		di.InitRepositories,
		di.InitServices,
		di.InitControllers,
		di.NewServer,
	)

	return &gin.Engine{}
}
