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
	eventBus event.Bus
	service  service.Depositer
}

func NewDeposit(s service.Depositer, b event.Bus) Deposit {
	return Deposit{service: s, eventBus: b}
}

func (ctrl Deposit) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req request.Deposit
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	userId := c.Get("userId").(string)
	if err := ctrl.service.Deposit(ctx, transaction.Transaction{
		Type:      transaction.Deposit,
		AccountId: req.Account,
		Amount:    req.Amount,
		UserId:    userId,
	}); err != nil {
		log.Printf("[Deposit:Handle]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "could not deposit amount into the account"})
	}

	if err := ctrl.publishEvents(ctx, ctrl.service.PullEvents()); err != nil {
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
