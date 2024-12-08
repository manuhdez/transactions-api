package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

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

func (ctrl *FindAccountTransactions) Handle(c echo.Context) error {
	accountId := c.Param("id")
	ctx := c.Request().Context()

	transactions, err := ctrl.repository.FindByAccount(ctx, accountId)
	if err != nil {
		log.Printf("Failed to get transactions: %e", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "failed to find transactions", "error": err})
	}

	trx := make([]infra.JsonTransaction, len(transactions))
	for i := range transactions {
		trx[i] = infra.NewJsonTransaction(transactions[i])
	}

	return c.JSON(http.StatusOK, response{trx})
}
