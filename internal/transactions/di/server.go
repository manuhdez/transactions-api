package di

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/http/router"
)

type Server struct {
	Engine   *gin.Engine
	EventBus event.Bus
}

func NewServer(
	eventBus event.Bus,
	router router.Router,
	accountCreatedHandler handler.AccountCreated,
) Server {
	// Register event handlers
	eventBus.Subscribe(event.AccountCreatedType, accountCreatedHandler)

	return Server{Engine: router.Engine, EventBus: eventBus}
}
