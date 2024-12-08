package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type transferRequest struct {
	From   string  `json:"from" validate:"required"`
	To     string  `json:"to" validate:"required"`
	Amount float32 `json:"amount" validate:"required,min=0"`
}

type Transfer struct {
	transferService service.Transfer
	accFinder       *service.AccountFinder
}

func NewTransferController(srv service.Transfer, af *service.AccountFinder) Transfer {
	return Transfer{
		transferService: srv,
		accFinder:       af,
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

	origin, err := ctrl.accFinder.Find(ctx, req.From)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "origin account not found"})
	}

	destination, err := ctrl.accFinder.Find(ctx, req.To)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "destination account not found"})
	}

	userId := c.Get("userId").(string)
	transfer := transaction.CreateTransfer(domain.NewID(userId), origin, destination, req.Amount)

	if err = ctrl.transferService.Transfer(ctx, transfer); err != nil {
		log.Printf("[Transfer:Handle][Transfer]%s", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"msg": "transfer operation failed"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Transfer finished successfully"})
}
