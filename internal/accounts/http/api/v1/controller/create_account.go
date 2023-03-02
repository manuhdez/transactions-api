package controller

import (
	"log"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/request"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type CreateAccount struct {
	service service.CreateService
}

func NewCreateAccount(s service.CreateService) CreateAccount {
	return CreateAccount{s}
}

func (c CreateAccount) Handle(ctx *gin.Context) {
	var req request.CreateAccount
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc := account.New(req.Id, req.Balance, req.Currency)
	err = c.service.Create(acc)
	if err != nil {
		log.Printf("Error creating account: %e", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": acc.Id()})
}
