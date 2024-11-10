package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
)

type DeleteAccount struct {
	service service.DeleteAccountService
}

func NewDeleteAccount(s service.DeleteAccountService) DeleteAccount {
	return DeleteAccount{s}
}

func (ctrl DeleteAccount) Handle(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	if err := ctrl.service.Delete(ctx, id); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusOK, echo.Map{"msg": "account deleted."})
}
