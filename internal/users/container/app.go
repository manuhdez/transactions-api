package container

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/config"
	domainservice "github.com/manuhdez/transactions-api/internal/users/domain/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/http/api"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

type App struct {
	Server api.Router
}

func NewApp(server api.Router) App {
	return App{Server: server}
}

var Databases = wire.NewSet(
	config.NewDBConnection,
)

var Repositories = wire.NewSet(
	wire.Bind(new(user.Repository), new(infra.UserMysqlRepository)),
	infra.NewUserMysqlRepository,
)

var DomainServices = wire.NewSet(
	wire.Bind(new(domainservice.HashService), new(infra.BcryptHashService)),
	infra.NewBcryptService,
)

var Services = wire.NewSet(
	service.NewRegisterUserService,
	service.NewLoginService,
	wire.Bind(new(infra.TokenService), new(infra.JWTService)),
	infra.NewJWTService,
)

var Controllers = wire.NewSet(
	controller.NewRegisterUserController,
	controller.NewLoginController,
)

var Router = wire.NewSet(api.NewRouter)
