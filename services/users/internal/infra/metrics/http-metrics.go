package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// requestDuration tracks the duration of HTTP requests in milliseconds
	requestDuration = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: "users",
		Name:      "http_request_duration_milliseconds_gauge",
		Help:      "Duration of HTTP requests in milliseconds",
	}, []string{"path"})

	// requestsCount tracks the total number of HTTP requests received
	requestsCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "users",
		Name:      "http_requests_total",
		Help:      "Total number of HTTP requests",
	}, []string{"method", "path"})

	// currentRequests tracks the current number of HTTP requests being handled
	currentRequests = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "users",
		Name:      "http_requests_in_flight",
		Help:      "Current number of HTTP requests being served",
	})
)

func TrackRequestDuration(since time.Time, path string) {
	requestDuration.WithLabelValues(path).Add(float64(time.Since(since).Milliseconds()))
}

func TrackRequestCountInc(method, pattern string) {
	requestsCount.WithLabelValues(method, pattern).Inc()
}

func TrackCurrentRequestsInc() {
	currentRequests.Inc()
}

func TrackCurrentRequestsDec() {
	currentRequests.Dec()
}
