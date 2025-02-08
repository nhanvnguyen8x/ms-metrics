package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP Requests total",
		},
		[]string{"method", "path"},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP Request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	UserRegistrations = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "user_registrations_total",
			Help: "User registrations total",
		},
	)

	ActiveConnections = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_connections",
			Help: "Active connections",
		},
	)

	RequestSizeHistogram = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "request_size_bytes",
			Help:    "Request size in bytes",
			Buckets: prometheus.ExponentialBuckets(100, 2, 10),
		},
	)
)

func init() {
	fmt.Println("Registered metrics")

	prometheus.MustRegister(requestsTotal, requestDuration)
	prometheus.MustRegister(UserRegistrations)
	prometheus.MustRegister(ActiveConnections)
	prometheus.MustRegister(RequestSizeHistogram)
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		duration := time.Since(start).Seconds()
		requestsTotal.WithLabelValues(c.Request.Method, c.FullPath()).Inc()
		requestDuration.WithLabelValues(c.Request.Method, c.FullPath()).Observe(duration)
	}
}
