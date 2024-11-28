package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/api"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
)

type findAllAccountsResponse struct {
	Accounts []api.AccountJson `json:"accounts"`
}

type FindAllAccounts struct {
	service service.AccountsFinder
}

func NewFindAllAccounts(s service.AccountsFinder) FindAllAccounts {
	return FindAllAccounts{s}
}

// Handle handles the request for retrieving a list of accounts
func (ctrl FindAllAccounts) Handle(c echo.Context) error {
	userId, ok := c.Get("userId").(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "you need to be logged in"})
	}

	ctx := c.Request().Context()
	accounts, err := ctrl.service.FindByUserId(ctx, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, findAllAccountsResponse{Accounts: api.NewJsonAccountList(accounts)})
}
