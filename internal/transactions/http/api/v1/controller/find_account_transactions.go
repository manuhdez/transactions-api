package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
)

type FindAccountTransactions struct {
	repository transaction.Repository
}

func NewFindAccountTransactions(repo transaction.Repository) FindAccountTransactions {
	return FindAccountTransactions{repo}
}

type response struct {
	Transactions []infra.JsonTransaction `json:"transactions"`
}

func (ctlr *FindAccountTransactions) Handle(ctx *gin.Context) {
	accountId := ctx.Param("id")
	transactions, err := ctlr.repository.FindByAccount(ctx, accountId)
	if err != nil {
		log.Printf("Failed to get transactions: %e", err)
		ctx.Status(http.StatusInternalServerError)
		return
	}

	tt := make([]infra.JsonTransaction, 0)
	for _, t := range transactions {
		tt = append(tt, infra.NewJsonTransaction(t))
	}

	res := response{tt}
	ctx.JSON(http.StatusOK, res)
}
