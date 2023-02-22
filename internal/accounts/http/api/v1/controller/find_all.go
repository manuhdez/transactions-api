package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

type FindAllAccounts struct {
	service service.FindAllService
}

func NewFindAllAccounts(s service.FindAllService) FindAllAccounts {
	return FindAllAccounts{s}
}

func (c FindAllAccounts) Handle(ctx *gin.Context) {
	accounts, err := c.service.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := infra.NewJsonAccountList(accounts)
	ctx.JSON(http.StatusOK, response)
}
