package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
)

func Server() *gin.Engine {
	deps := Deps()

	server := gin.Default()
	server.GET("/status", controllers.StatusController)
	server.POST("/accounts", controllers.CreateAccountController(deps.Services.CreateAccount))
	server.GET("/accounts", controllers.FindAllController(deps.Services.FindAll))

	return server
}
