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

func CreateAccountController(service service.CreateService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createAccountRequest
		ctx.BindJSON(&req)

		acc := account.New(req.Id, req.Balance)
		err := service.Create(acc)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
