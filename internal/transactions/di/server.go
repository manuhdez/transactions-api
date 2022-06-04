package di

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
)

func NewServer(
	statusController controllers.StatusController,
) *gin.Engine {
	srv := gin.Default()

	// Register server routes
	srv.GET("/status", statusController.Handle)

	return srv
}
