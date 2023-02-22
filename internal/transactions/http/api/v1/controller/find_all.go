package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

type FindAllTransactions struct {
	service service.FindAllTransactions
}

func NewFindAllTransactions(s service.FindAllTransactions) FindAllTransactions {
	return FindAllTransactions{s}
}

func (c FindAllTransactions) Handle(ctx *gin.Context) {
	transactions, err := c.service.Invoke(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error fetching the list of transactions"})
	}

	ctx.JSON(http.StatusOK, transactions)
}
