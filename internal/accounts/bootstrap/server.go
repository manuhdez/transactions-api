package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/handler"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/http/router"
)

type Server struct {
	Engine   *gin.Engine
	EventBus event.Bus
}

func InitServer(
	eventBus event.Bus,
	router router.Router,
	depositCreatedHandler handler.DepositCreated,
	withdrawCreatedHandler handler.WithdrawCreated,
) Server {
	eventBus.Subscribe(handler.DepositCreatedType, depositCreatedHandler)
	eventBus.Subscribe(handler.WithdrawCreatedType, withdrawCreatedHandler)

	return Server{Engine: router.Engine, EventBus: eventBus}
}
