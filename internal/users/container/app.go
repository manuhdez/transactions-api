package container

import (
	"github.com/google/wire"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/config"
	"github.com/manuhdez/transactions-api/internal/users/domain/event"
	domainservice "github.com/manuhdez/transactions-api/internal/users/domain/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/http/api"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

type App struct {
	Server   api.Router
	EventBus event.Bus
}

func NewApp(server api.Router, eventBus event.Bus) App {
	return App{
		Server:   server,
		EventBus: eventBus,
	}
}

var Databases = wire.NewSet(
	config.NewDBConnection,
)

var Buses = wire.NewSet(
	wire.Bind(new(event.Bus), new(infra.RabbitEventBus)),
	infra.NewRabbitEventBus,
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
	service.NewUsersRetrieverService,
)

var Controllers = wire.NewSet(
	controller.NewHealthCheck,
	controller.NewRegisterUserController,
	controller.NewLoginController,
	controller.NewGetAllUsersController,
)

var Router = wire.NewSet(api.NewRouter)
