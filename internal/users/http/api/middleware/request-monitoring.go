package middleware

import (
	"net/http"
	"time"

	"github.com/manuhdez/transactions-api/internal/users/infra/metrics"
)

func RequestMonitoring(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Track request duration
		defer metrics.TrackRequestDuration(time.Now(), r.URL.Path)

		// Track in-flight requests
		metrics.TrackCurrentRequestsInc()
		defer metrics.TrackCurrentRequestsDec()

		// Increment total requests
		metrics.TrackRequestCountInc(r.Method, r.Pattern)

		// Call next handler
		next.ServeHTTP(w, r)
	})
}
