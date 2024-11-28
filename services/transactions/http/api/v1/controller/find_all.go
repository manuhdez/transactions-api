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
	Error        string                    `json:"error,omitempty"`
}

type FindAllTransactions struct {
	finder service.TransactionsRetriever
}

func NewFindAllTransactions(s service.TransactionsRetriever) FindAllTransactions {
	return FindAllTransactions{s}
}

func (ctrl FindAllTransactions) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get("userId").(string)

	transactions, err := ctrl.finder.Retrieve(ctx, userId)
	if err != nil {
		log.Printf("[TransactionsRetriever:Handle]%s", err)
		return c.JSON(
			http.StatusInternalServerError,
			findAllTransactionsResponse{Error: "failed to fetch transactions"},
		)
	}

	return c.JSON(http.StatusOK, findAllTransactionsResponse{Transactions: transactions})
}
