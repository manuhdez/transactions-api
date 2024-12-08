package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type withdrawRequest struct {
	Account  string  `json:"account" validate:"required"`
	Amount   float32 `json:"amount" validate:"required"`
	Currency string  `json:"currency" validate:"required,iso4217"`
}

type Withdraw struct {
	withdrawSrv *service.WithdrawService
	accFinder   *service.AccountFinder
}

func NewWithdraw(s *service.WithdrawService, af *service.AccountFinder) Withdraw {
	return Withdraw{
		withdrawSrv: s,
		accFinder:   af,
	}
}

func (ctrl Withdraw) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req withdrawRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("[Withdraw:Handle][Bind]%s", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "Missing params", "error": err})
	}

	if err := c.Validate(req); err != nil {
		log.Printf("[Withdraw:Handle][Validate]%s", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "invalid request data"})
	}

	userId := c.Get("userId").(string)

	// Check if user has access to account
	account, err := ctrl.accFinder.Find(ctx, req.Account)
	if err != nil || account.UserId.String() != userId {
		fmt.Printf("[Withdraw:Handle]%s", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	trx := transaction.CreateWithdrawal(account, domain.NewID(userId), req.Amount)
	if err := ctrl.withdrawSrv.Withdraw(ctx, trx); err != nil {
		log.Printf("[Withdraw:Handle][Withdraw]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "withdraw operation failed"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Withdraw successfully created"})
}
