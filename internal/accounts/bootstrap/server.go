package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/handler"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/http/middleware"
)

type Server struct {
	Engine   *gin.Engine
	EventBus event.Bus
}

func InitServer(
	eventBus event.Bus,
	depositCreatedHandler handler.DepositCreated,
	withdrawCreatedHandler handler.WithdrawCreated,
	status controllers.StatusController,
	findAll controllers.FindAllAccountsController,
	create controllers.CreateAccountController,
	find controllers.FindAccountController,
	deleteAccount controllers.DeleteAccountController,
) Server {
	engine := gin.Default()

	// Register global middleware
	engine.Use(middleware.CORSMiddleware())

	engine.GET("/status", status.Handle)
	engine.GET("/accounts", findAll.Handle)
	engine.POST("/accounts", create.Handle)
	engine.GET("/accounts/:id", find.Handle)
	engine.DELETE("/accounts/:id", deleteAccount.Handle)

	eventBus.Subscribe(handler.DepositCreatedType, depositCreatedHandler)
	eventBus.Subscribe(handler.WithdrawCreatedType, withdrawCreatedHandler)

	return Server{Engine: engine, EventBus: eventBus}
}
