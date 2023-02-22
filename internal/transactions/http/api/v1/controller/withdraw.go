package controller

import (
	"context"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

type Withdraw struct {
	service service.Withdraw
}

func NewWithdraw(s service.Withdraw) Withdraw {
	return Withdraw{s}
}

func (c Withdraw) Handle(ctx *gin.Context) {
	var req request.Withdraw
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing params"})
	}

	withdraw := transaction.NewTransaction(transaction.Withdrawal, req.Account, req.Amount, req.Currency)
	if err := c.service.Invoke(context.Background(), withdraw); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create withdraw"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Withdraw successfully created"})
}
