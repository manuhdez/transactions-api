//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/di"
)

func InitServer() *http.ServeMux {
	wire.Build(di.InitControllers, di.NewServer)
	return &http.ServeMux{}
}
