package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/request"
)

var (
	ErrInvalidUserIdForCreate = errors.New("cannot create an account for a different user than the logged in")
	ErrUnauthorized           = errors.New("unauthorized to create account")
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
		log.Printf("[CreateAccount:Handle][err: %s]", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err})
	}

	userId, ok := c.Get("userId").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": ErrUnauthorized})
	}

	ctx := c.Request().Context()
	if err := ctrl.service.Create(ctx, req.Decode(userId)); err != nil {
		log.Printf("Error creating account: %e", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Account created successfully!"})
}
