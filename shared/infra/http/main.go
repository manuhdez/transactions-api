package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func EchoWrapper(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
