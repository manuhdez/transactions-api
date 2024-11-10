package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthCheck struct{}

func NewHealthCheck() HealthCheck {
	return HealthCheck{}
}

func (ctrl HealthCheck) Handle(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"status": "ok", "service": "users"})
}
