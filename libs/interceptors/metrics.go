package interceptors

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	RequestsCounter       *prometheus.CounterVec
	ResponseCounter       *prometheus.CounterVec
	HistogramResponseTime *prometheus.HistogramVec
)

func InitMetricsInterceptor(serviceName string) {
	RequestsCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: "grpc",
		Name:      "requests_total",
	}, []string{"handler"})
	ResponseCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: "grpc",
		Name:      "responses_total",
	}, []string{"handler", "status"})
	HistogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: serviceName,
		Subsystem: "grpc",
		Name:      "histogram_response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	}, []string{"handler", "status"})
}

func MetricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	method := info.FullMethod
	RequestsCounter.WithLabelValues(method).Inc()
	start := time.Now()
	res, err := handler(ctx, req)
	code := status.Convert(err).Code().String()

	if err != nil {
		ResponseCounter.WithLabelValues(method, code).Inc()
		return nil, err
	}

	elapsed := time.Since(start)
	HistogramResponseTime.WithLabelValues(method, code).Observe(elapsed.Seconds())
	ResponseCounter.WithLabelValues(method, code).Inc()
	return res, nil
}
