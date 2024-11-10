package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

type FindAllAccounts struct {
	service service.FindAllService
}

func NewFindAllAccounts(s service.FindAllService) FindAllAccounts {
	return FindAllAccounts{s}
}

func (ctrl FindAllAccounts) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	accounts, err := ctrl.service.Find(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	response := infra.NewJsonAccountList(accounts)
	return c.JSON(http.StatusOK, response)
}
