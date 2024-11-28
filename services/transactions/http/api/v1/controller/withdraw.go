package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
)

type Withdraw struct {
	eventBus   event.Bus
	trxService *service.TransactionService
	accFinder  *service.AccountFinder
}

func NewWithdraw(s *service.TransactionService, af *service.AccountFinder, b event.Bus) Withdraw {
	return Withdraw{
		trxService: s,
		accFinder:  af,
		eventBus:   b,
	}
}

func (ctrl Withdraw) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.Withdraw
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
	if err != nil || account.UserId != userId {
		fmt.Printf("[Withdraw:Handle]%s", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	if err := ctrl.trxService.Withdraw(ctx, transaction.Transaction{
		Type:      transaction.Withdrawal,
		AccountId: req.Account,
		Amount:    req.Amount,
		UserId:    userId,
	}); err != nil {
		log.Printf("[Withdraw:Handle][Withdraw]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "withdraw operation failed"})
	}

	if err := ctrl.publishEvents(ctx, ctrl.trxService.PullEvents()); err != nil {
		log.Printf("[Withdraw:Handle]%s", err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Withdraw successfully created"})
}

func (ctrl Withdraw) publishEvents(ctx context.Context, events []event.Event) error {
	var errList []error
	for i := range events {
		if err := ctrl.eventBus.Publish(ctx, events[i]); err != nil {
			errList = append(errList, err)
		}
	}

	if len(errList) > 0 {
		return fmt.Errorf("[publishEvents][err: failed to publish %d events][errors %+v]", len(errList), errList)
	}

	return nil
}
