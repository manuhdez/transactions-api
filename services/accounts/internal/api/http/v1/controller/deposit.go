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

type depositRequest struct {
	Account  string  `json:"account" binding:"required"`
	Amount   float32 `json:"amount" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
}

type Deposit struct {
	depositService *service.DepositService
	accountFinder  *service.AccountFinder
}

func NewDeposit(s *service.DepositService, f *service.AccountFinder) Deposit {
	return Deposit{
		depositService: s,
		accountFinder:  f,
	}
}

func (ctrl Deposit) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req depositRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	userId := c.Get("userId").(string)

	// Check if user has access to account
	account, err := ctrl.accountFinder.Find(ctx, req.Account)
	if err != nil || account.UserId.String() != userId {
		fmt.Printf("[Deposit:Handle]%s", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	// Execute deposit transaction
	deposit := transaction.CreateDeposit(account, domain.NewID(userId), req.Amount)
	if err = ctrl.depositService.Deposit(ctx, deposit); err != nil {
		log.Printf("[Deposit:Handle]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "could not deposit amount into the account"})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Deposit successfully created"})
}
