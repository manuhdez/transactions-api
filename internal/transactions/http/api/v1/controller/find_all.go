package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type findAllTransactionsResponse struct {
	Transactions []transaction.Transaction `json:"transactions"`
}

type FindAllTransactions struct {
	service service.FindAllTransactions
}

func NewFindAllTransactions(s service.FindAllTransactions) FindAllTransactions {
	return FindAllTransactions{s}
}

func (ctrl FindAllTransactions) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	transactions, err := ctrl.service.Invoke(ctx)
	if err != nil {
		log.Printf("[FindAllTransactions:Handle]%s", err)
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"message": "Could not retrieve transactions"},
		)
	}

	return c.JSON(http.StatusOK, findAllTransactionsResponse{Transactions: transactions})
}
