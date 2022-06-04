package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type createAccountRequest struct {
	Id      string  `json:"id"`
	Balance float32 `json:"balance"`
}

type CreateAccountController struct {
	service service.CreateService
}

func NewCreateAccountController(s service.CreateService) CreateAccountController {
	return CreateAccountController{s}
}

func (c CreateAccountController) Handle(ctx *gin.Context) {
	var req createAccountRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	acc := account.New(req.Id, req.Balance)
	err = c.service.Create(acc)
	if err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"id": acc.Id()})
}
