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

type Deposit struct {
	eventBus      event.Bus
	trxService    *service.TransactionService
	accountFinder *service.AccountFinder
}

func NewDeposit(s *service.TransactionService, f *service.AccountFinder, b event.Bus) Deposit {
	return Deposit{
		trxService:    s,
		accountFinder: f,
		eventBus:      b,
	}
}

func (ctrl Deposit) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.Deposit
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	userId := c.Get("userId").(string)

	// Check if user has access to account
	account, err := ctrl.accountFinder.Find(ctx, req.Account)
	if err != nil || account.UserId != userId {
		fmt.Printf("[Deposit:Handle]%s", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthorized"})
	}

	// Execute deposit transaction
	deposit := transaction.NewDeposit(req.Account, userId, req.Amount)
	if err = ctrl.trxService.Deposit(ctx, deposit); err != nil {
		log.Printf("[Deposit:Handle]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "could not deposit amount into the account"})
	}

	if err = ctrl.publishEvents(ctx, ctrl.trxService.PullEvents()); err != nil {
		log.Printf("[Deposit:Handle]%s", err)
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Deposit successfully created"})
}

func (ctrl Deposit) publishEvents(ctx context.Context, events []event.Event) error {
	var errList []error
	for i := range events {
		if err := ctrl.eventBus.Publish(ctx, events[i]); err != nil {
			log.Printf("[Deposit:publishEvents][failed to publish event]%s", err)
			errList = append(errList, err)
		}
	}

	if len(errList) > 0 {
		return fmt.Errorf("[publishEvents][err: failed to publish %d events][errors: %+v]", len(errList), errList)
	}

	return nil
}
