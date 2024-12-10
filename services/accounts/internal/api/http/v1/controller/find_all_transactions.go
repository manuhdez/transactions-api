package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/infra/db"
)

type findAllTransactionsResponse struct {
	Transactions []db.JsonTransaction `json:"transactions"`
	Error        string               `json:"error,omitempty"`
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

	return c.JSON(http.StatusOK, newFindAllTransactionsResponse(transactions))
}

func newFindAllTransactionsResponse(trx []transaction.Transaction) findAllTransactionsResponse {
	transactions := make([]db.JsonTransaction, len(trx))
	for i := range trx {
		transactions[i] = db.NewJsonTransaction(trx[i])
	}

	return findAllTransactionsResponse{Transactions: transactions}
}
