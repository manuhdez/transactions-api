package bootstrap

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func InitializeServer(c Controllers) Server {
	engine := Engine(c)
	return Server{engine}
}

func Engine(c Controllers) *gin.Engine {
	server := gin.Default()
	server.GET("/status", c.Status.Handle)
	server.GET("/accounts", c.FindAllAccounts.Handle)
	server.POST("/accounts", c.CreateAccount.Handle)
	server.GET("/accounts/:id", c.FindAccount.Handle)
	return server
}
