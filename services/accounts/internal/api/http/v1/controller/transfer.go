package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type transferRequest struct {
	From   string  `json:"from" validate:"required"`
	To     string  `json:"to" validate:"required"`
	Amount float32 `json:"amount" validate:"required"`
}

type Transfer struct {
	eventBus        event.Bus
	transferService *service.TransactionService
	accFinder       *service.AccountFinder
}

func NewTransferController(srv *service.TransactionService, af *service.AccountFinder, bus event.Bus) Transfer {
	return Transfer{
		transferService: srv,
		accFinder:       af,
		eventBus:        bus,
	}
}

func (ctrl Transfer) Handle(c echo.Context) error {
	ctx := c.Request().Context()

	var req transferRequest
	if err := c.Bind(&req); err != nil {
		log.Printf("[Transfer:Handle][Bind]%s", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "failed to parse request"})
	}

	if err := c.Validate(req); err != nil {
		log.Printf("[Transfer:Handle][Validate]%s", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"msg": "invalid request data", "error": err})
	}

	// Check accounts exist
	origin, err := ctrl.accFinder.Find(ctx, req.From)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "origin account not found"})
	}
	_, err = ctrl.accFinder.Find(ctx, req.To)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "destination account not found"})
	}

	// Check the transaction can be made
	if ok := ctrl.isTransferAllowed(c, req, origin); !ok {
		log.Printf("[Transfer:Handle][CanTransfer]%s", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{"msg": "unauthorized request", "error": err})
	}

	userId := c.Get("userId").(string)
	if err = ctrl.transferService.Transfer(ctx, transaction.NewTransfer(userId, req.From, req.To, req.Amount)); err != nil {
		log.Printf("[Transfer:Handle][Transfer]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "transfer operation failed"})
	}

	if err := ctrl.publishEvents(ctx); err != nil {
		log.Printf("[Transfer:Handle]%s", err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Transfer finished successfully"})
}

// isTransferAllowed checks if the user has access to the account and it has enough balance to perform the transfer
func (ctrl Transfer) isTransferAllowed(c echo.Context, req transferRequest, origin account.Account) bool {
	userId := c.Get("userId").(string)

	if origin.UserId.String() != userId {
		log.Printf("[isTransferAllowed][userId: %s][msg: user does not have access to the origin account]", userId)
		return false
	}

	// TODO: check account balance before transfer
	// if origin.Balance < req.Amount {
	// 	log.Printf("[isTransferAllowed][originAccount:%+v][err: not enough balance]", origin)
	// 	return false, nil
	// }

	return true
}

// publishEvents publish the
func (ctrl Transfer) publishEvents(ctx context.Context) error {
	var errorList []error

	events := ctrl.transferService.PullEvents()
	if len(events) == 0 {
		return fmt.Errorf("[publishEvents][err: no events where generated]")
	}

	for i := range events {
		if err := ctrl.eventBus.Publish(ctx, events[i]); err != nil {
			// TODO: handle retry of unpublished events
			errorList = append(errorList, err)
		} else {
			log.Printf("[Transfer][publishEvents][msg: new event published][event: %+v]", events[i])
		}
	}

	if len(errorList) > 0 {
		return fmt.Errorf("[publishEvents][err: failed to publish %d events]", len(errorList))
	}

	return nil
}
