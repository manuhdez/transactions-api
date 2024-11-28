package controller

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	domainshared "github.com/manuhdez/transactions-api/shared/domain"
)

var ErrUnauthorized = errors.New("unauthorized to create account")

type createAccountRequest struct {
	Id       string  `json:"id" validate:"required"`
	Balance  float32 `json:"balance" default:"0"`
	Currency string  `json:"currency" default:"EUR"`
}

// decode turns a createAccountRequest into a domain Account object
func (req createAccountRequest) decode(userId string) account.Account {
	return account.NewWithUserID(req.Id, domainshared.NewID(userId), req.Balance, req.Currency)
}

type CreateAccount struct {
	service service.CreateService
}

func NewCreateAccount(s service.CreateService) CreateAccount {
	return CreateAccount{s}
}

func (ctrl CreateAccount) Handle(c echo.Context) error {
	var req createAccountRequest
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
	if err := ctrl.service.Create(ctx, req.decode(userId)); err != nil {
		log.Printf("Error creating account: %e", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "Account created successfully!"})
}
