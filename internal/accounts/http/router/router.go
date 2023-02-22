package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
	"github.com/manuhdez/transactions-api/internal/accounts/http/middleware"
	"net/http"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(
	findAllAccounts controllers.FindAllAccountsController,
	createAccount controllers.CreateAccountController,
	findAccount controllers.FindAccountController,
	deleteAccount controllers.DeleteAccountController,
) Router {
	router := gin.Default()

	// Register global middleware
	router.Use(middleware.CORSMiddleware())

	router.GET("/status", statusHandler)

	api := router.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/accounts", findAllAccounts.Handle)
		v1.POST("/accounts", createAccount.Handle)
		v1.GET("/accounts/:id", findAccount.Handle)
		v1.DELETE("/accounts/:id", deleteAccount.Handle)
	}

	return Router{router}
}

func statusHandler(ctx *gin.Context) {
	type statusResponse struct {
		Status  string `json:"status"`
		Service string `json:"service"`
	}
	ctx.JSON(http.StatusOK, statusResponse{"ok", "accounts"})
}
