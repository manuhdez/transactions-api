package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/accounts/http/middleware"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter(
	findAllAccounts controller.FindAllAccounts,
	createAccount controller.CreateAccount,
	findAccount controller.FindAccount,
	deleteAccount controller.DeleteAccount,
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
