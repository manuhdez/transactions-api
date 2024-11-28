package middleware

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/users/internal/infra/metrics"
)

func RequestMonitoring(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()

		// Track request duration
		defer metrics.TrackRequestDuration(time.Now(), r.URL.Path)

		// Track in-flight requests
		metrics.TrackCurrentRequestsInc()
		defer metrics.TrackCurrentRequestsDec()

		// Increment total requests
		metrics.TrackRequestCountInc(r.Method, r.Pattern)

		// Call next handler
		return next(c)
	}
}
