package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

type FindAllAccountsController struct {
	service service.FindAllService
}

func NewFindAllAccountsControllers(s service.FindAllService) FindAllAccountsController {
	return FindAllAccountsController{s}
}

func (c FindAllAccountsController) Handle(ctx *gin.Context) {
	accounts, err := c.service.Find(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	response := infra.NewJsonAccountList(accounts)
	ctx.JSON(http.StatusOK, response)
}
