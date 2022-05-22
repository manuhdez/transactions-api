package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

func FindAllController(s service.FindAllService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accounts, err := s.Find(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		response := infra.NewJsonAccountList(accounts)
		ctx.JSON(http.StatusOK, response)
	}
}
