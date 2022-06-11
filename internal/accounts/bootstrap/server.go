package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
)

type Server struct {
	Engine *gin.Engine
}

func InitServer(
	status controllers.StatusController,
	findAll controllers.FindAllAccountsController,
	create controllers.CreateAccountController,
	find controllers.FindAccountController,
	deleteAccount controllers.DeleteAccountController,
) Server {
	engine := gin.Default()
	engine.GET("/status", status.Handle)
	engine.GET("/accounts", findAll.Handle)
	engine.POST("/accounts", create.Handle)
	engine.GET("/accounts/:id", find.Handle)
	engine.DELETE("/accounts/:id", deleteAccount.Handle)
	return Server{Engine: engine}
}
