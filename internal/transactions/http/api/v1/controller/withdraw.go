package controller

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

type Withdraw struct {
	service service.Withdraw
}

func NewWithdraw(s service.Withdraw) Withdraw {
	return Withdraw{s}
}

func (ctrl Withdraw) Handle(c echo.Context) error {
	var req request.Withdraw
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "Missing params", "error": err})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	withdraw := transaction.NewTransaction(transaction.Withdrawal, req.Account, req.Amount, req.Currency)
	if err := ctrl.service.Invoke(context.Background(), withdraw); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "could not create withdraw", "error": err})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Withdraw successfully created"})
}
