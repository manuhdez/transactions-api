package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
)

type Deposit struct {
	service service.Deposit
}

func NewDeposit(s service.Deposit) Deposit {
	return Deposit{s}
}

func (c Deposit) Handle(ctx *gin.Context) {
	var req request.Deposit
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deposit := transaction.NewTransaction(transaction.Deposit, req.Account, req.Amount, req.Currency)

	err := c.service.Invoke(ctx, deposit)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Deposit successfully created"})
}
