package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type DepositController struct {
	service service.Deposit
}

func NewDepositController(s service.Deposit) DepositController {
	return DepositController{s}
}

type DepositRequest struct {
	Account  string  `json:"account"`
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

func (c DepositController) Handle(ctx *gin.Context) {
	var req DepositRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	deposit := transaction.NewTransaction(transaction.Deposit, req.Amount, req.Currency)

	err := c.service.Invoke(ctx, deposit)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Deposit successfully created"})
}
