package inmemory

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	CacheHitsTotal             prometheus.Counter
	CacheErrorsTotal           prometheus.Counter
	CacheRequestsTotal         prometheus.Counter
	HistogramResponseTimeCache prometheus.Histogram
)

func InitMetricsCache(serviceName string) {
	CacheHitsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: "cache",
		Name:      "hits_total",
	})
	CacheErrorsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: "cache",
		Name:      "errors_total",
	})
	CacheRequestsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: "cache",
		Name:      "requests_total",
	})
	HistogramResponseTimeCache = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: serviceName,
		Subsystem: "cache",
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.000001, 2, 16),
	})
}
