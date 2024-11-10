package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

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
		return c.JSON(
			http.StatusInternalServerError,
			echo.Map{"message": "There was an error fetching the list of transactions", "error": err},
		)
	}

	return c.JSON(http.StatusOK, transactions)
}
