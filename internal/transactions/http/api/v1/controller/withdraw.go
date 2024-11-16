package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
)

type Withdraw struct {
	trxService service.TransactionService
}

func NewWithdraw(s service.TransactionService) Withdraw {
	return Withdraw{trxService: s}
}

func (ctrl Withdraw) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.Withdraw
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "Missing params", "error": err})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	userId := c.Get("userId").(string)
	if err := ctrl.trxService.Withdraw(ctx, transaction.Transaction{
		Type:      transaction.Withdrawal,
		AccountId: req.Account,
		Amount:    req.Amount,
		UserId:    userId,
	}); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "could not create withdraw", "error": err})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Withdraw successfully created"})
}
