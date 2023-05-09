package observability

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	httpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path"},
	)

	activeRequests = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_requests",
			Help: "Number of active HTTP requests.",
		},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "The HTTP request duration in seconds.",
			Buckets: []float64{0.1, 0.25, 0.5, 1, 2.5, 5, 10},
		},
		// requests that take less than 0.1 seconds
		// requests that take between 0.1 and 0.25 seconds
		// requests that take between 0.25 and 0.5 seconds
		// requests that take between 0.5 and 1 second
		// requests that take between 1 and 2.5 seconds
		// requests that take between 2.5 and 5 seconds
		// requests that take between 5 and 10 seconds.
		[]string{"method", "path", "status_code"},
	)
)

func InitMetrics(router *gin.Engine) {
	// Unregister Previous Metrics
	prometheus.Unregister(httpRequestsTotal)
	prometheus.Unregister(activeRequests)
	prometheus.Unregister(httpRequestDuration)

	// Register the Prometheus metrics collectors
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(activeRequests)
	prometheus.MustRegister(httpRequestDuration)
}

func CalcRequests(ctx *gin.Context) {
	activeRequests.Inc()

	// Call SetMetrics() after the request is complete
	defer SetMetrics(ctx)

	ctx.Next()
	activeRequests.Dec()
}

func CalcRequestDurations(ctx *gin.Context, start time.Time) {
	activeRequests.Inc()

	// Call SetDuration() after the request is complete
	defer SetDuration(ctx, start)

	ctx.Next()
	activeRequests.Dec()
}

func SetMetrics(ctx *gin.Context) {
	httpRequestsTotal.With(prometheus.Labels{
		"method": ctx.Request.Method,
		"path":   ctx.Request.URL.Path,
	}).Inc()
}

func SetDuration(ctx *gin.Context, start time.Time) {
	httpRequestDuration.With(prometheus.Labels{
		"method":      ctx.Request.Method,
		"path":        ctx.Request.URL.Path,
		"status_code": strconv.Itoa(ctx.Writer.Status()),
	}).Observe(float64(time.Since(start).Seconds()))
}
