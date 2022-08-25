package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

type FindAllTransactionsController struct {
	service service.FindAllTransactions
}

func NewFindAllController(s service.FindAllTransactions) FindAllTransactionsController {
	return FindAllTransactionsController{s}
}

func (c FindAllTransactionsController) Handle(ctx *gin.Context) {
	transactions, err := c.service.Invoke(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error fetching the list of transactions"})
	}

	ctx.JSON(http.StatusOK, transactions)
}
