package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/request"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type CreateAccount struct {
	service service.CreateService
}

func NewCreateAccount(s service.CreateService) CreateAccount {
	return CreateAccount{s}
}

func (ctrl CreateAccount) Handle(c echo.Context) error {
	var req request.CreateAccount
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	acc := account.New(req.Id, req.Balance, req.Currency)
	if err := ctrl.service.Create(acc); err != nil {
		log.Printf("Error creating account: %e", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusCreated, echo.Map{"id": acc.Id()})
}
