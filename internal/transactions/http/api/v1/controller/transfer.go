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
)

type transferRequest struct {
	From   string  `json:"from" validate:"required"`
	To     string  `json:"to" validate:"required"`
	Amount float32 `json:"amount" validate:"required"`
	UserId string  `json:"userId" validate:"required"`
}

type Transfer struct {
	eventBus        event.Bus
	transferService *service.TransactionService
}

func NewTransferController(srv *service.TransactionService, bus event.Bus) Transfer {
	return Transfer{transferService: srv, eventBus: bus}
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

	if err := ctrl.transferService.Transfer(ctx, transaction.NewTransfer(req.UserId, req.From, req.To, req.Amount)); err != nil {
		log.Printf("[Transfer:Handle][Transfer]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "transfer operation failed"})
	}

	if err := ctrl.publishEvents(ctx); err != nil {
		log.Printf("[Transfer:Handle]%s", err)
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Transfer finished successfully"})
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
