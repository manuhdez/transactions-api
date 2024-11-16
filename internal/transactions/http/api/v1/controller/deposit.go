package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
)

type Deposit struct {
	trxService service.TransactionService
}

func NewDeposit(s service.TransactionService) Deposit {
	return Deposit{trxService: s}
}

func (ctrl Deposit) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.Deposit
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	userId := c.Get("userId").(string)
	if err := ctrl.trxService.Deposit(ctx, transaction.Transaction{
		Type:      transaction.Deposit,
		AccountId: req.Account,
		Amount:    req.Amount,
		UserId:    userId,
	}); err != nil {
		log.Printf("[Deposit:Handle]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "could not deposit amount into the account"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Deposit successfully created"})
}
