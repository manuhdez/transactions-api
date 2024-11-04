package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Database metrics
var (
	// dbErrorsCount is a counter that tracks the total number of database errors thrown
	dbErrorsCount = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "users",
		Name:      "db_errors_count",
		Help:      "Total number of errors in database",
	})

	// dbQueryTime is a histogram that tracks the duration of database queries in milliseconds
	dbQueryTime = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "users",
		Name:      "db_query_response_time",
		Help:      "Duration of database queries in seconds",
		Buckets:   []float64{0.300, 0.500, 0.700, 1},
	})
)

func TrackDBErrorAdd() {
	dbErrorsCount.Inc()
}

func TrackDBQueryDuration(since time.Time) {
	dbQueryTime.Observe(time.Since(since).Seconds())
}
