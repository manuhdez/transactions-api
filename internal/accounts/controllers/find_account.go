package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

func FindAccountController(s service.FindAccountService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		acc, err := s.Find(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if (account.Account{} == acc) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}

		response := infra.NewJsonAccount(acc)
		ctx.JSON(http.StatusOK, response)
	}
}
