package controllers

import (
	"context"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

type WithdrawController struct {
	service service.Withdraw
}

func NewWithdrawController(s service.Withdraw) WithdrawController {
	return WithdrawController{s}
}

type WithdrawRequest struct {
	Account  string  `json:"account" binding:"required"`
	Amount   float32 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

func (c WithdrawController) Handle(ctx *gin.Context) {
	var req WithdrawRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing params"})
	}

	withdraw := transaction.NewTransaction(transaction.Withdrawal, req.Account, req.Amount, req.Currency)
	if err := c.service.Invoke(context.Background(), withdraw); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create withdraw"})
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Withdraw successfully created"})
}
