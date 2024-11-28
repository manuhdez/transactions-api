package main

import (
	"github.com/google/wire"

	"github.com/manuhdez/transactions-api/internal/users/internal/api/http/router"
	"github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/users/internal/domain/event"
	domainservice "github.com/manuhdez/transactions-api/internal/users/internal/domain/service"
	"github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/internal/infra"
	"github.com/manuhdez/transactions-api/internal/users/internal/infra/db"
	"github.com/manuhdez/transactions-api/internal/users/internal/infra/queue/rabbitmq"
	"github.com/manuhdez/transactions-api/shared/config"
)

type App struct {
	Server   router.Router
	EventBus event.Bus
}

func NewApp(server router.Router, eventBus event.Bus) App {
	return App{
		Server:   server,
		EventBus: eventBus,
	}
}

var Databases = wire.NewSet(
	config.NewDBConfig,
	config.NewGormDBConnection,
)

var Buses = wire.NewSet(
	wire.Bind(new(event.Bus), new(rabbitmq.RabbitEventBus)),
	rabbitmq.NewRabbitEventBus,
)

var Repositories = wire.NewSet(
	wire.Bind(new(user.Repository), new(db.UserPostgresRepository)),
	db.NewUserPostgresRepository,
)

var DomainServices = wire.NewSet(
	wire.Bind(new(domainservice.HashService), new(infra.BcryptHashService)),
	infra.NewBcryptService,
)

var Services = wire.NewSet(
	service.NewRegisterUserService,
	service.NewLoginService,
	wire.Bind(new(domainservice.TokenService), new(infra.JWTService)),
	infra.NewJWTService,
	service.NewUsersRetrieverService,
)

var Controllers = wire.NewSet(
	controller.NewHealthCheck,
	controller.NewRegisterUserController,
	controller.NewLoginController,
	controller.NewGetAllUsersController,
)

var Router = wire.NewSet(router.NewRouter)
