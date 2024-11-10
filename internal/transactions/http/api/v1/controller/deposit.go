package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
)

type Deposit struct {
	service service.Deposit
}

func NewDeposit(s service.Deposit) Deposit {
	return Deposit{s}
}

func (ctrl Deposit) Handle(c echo.Context) error {
	var req request.Deposit
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	deposit := transaction.NewTransaction(transaction.Deposit, req.Account, req.Amount, req.Currency)

	ctx := c.Request().Context()
	if err := ctrl.service.Invoke(ctx, deposit); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Deposit successfully created"})
}
