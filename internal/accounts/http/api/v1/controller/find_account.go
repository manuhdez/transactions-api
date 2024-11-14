package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

type FindAccount struct {
	service service.AccountsFinder
}

func NewFindAccountController(s service.AccountsFinder) FindAccount {
	return FindAccount{s}
}

func (ctrl FindAccount) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	acc, err := ctrl.service.FindById(ctx, c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	if (account.Account{} == acc) {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "account not found"})
	}

	return c.JSON(http.StatusOK, infra.NewJsonAccount(acc))
}
