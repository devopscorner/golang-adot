package observability

import (
	"strconv"
	"time"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/codes"
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
		[]string{"method", "path", "status_code"},
	)
)

func InitMetrics(router *gin.Engine) {
	// Register the Prometheus metrics collectors
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(activeRequests)
	prometheus.MustRegister(httpRequestDuration)

	// Set up the Prometheus middleware
	router.Use(PrometheusMiddleware())
}

func PrometheusMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start tracing
		_, span := StartTracing(ctx.Request.Context(), config.OtelServiceName())
		defer EndTracing(span, codes.Ok, config.OtelServiceName())

		SetRequests(ctx)
	}
}

func SetMetrics(ctx *gin.Context) {
	httpRequestsTotal.With(prometheus.Labels{
		"method": ctx.Request.Method,
		"path":   ctx.Request.URL.Path,
	}).Inc()
}

func SetRequests(ctx *gin.Context) {
	activeRequests.Inc()

	// Call SetMetrics() after the request is complete
	defer SetMetrics(ctx)

	// Call SetDuration() after the request is complete
	defer SetDuration(ctx, time.Now())

	ctx.Next()

	activeRequests.Dec()
}

func SetDuration(ctx *gin.Context, start time.Time) {
	httpRequestDuration.With(prometheus.Labels{
		"method":      ctx.Request.Method,
		"path":        ctx.Request.URL.Path,
		"status_code": strconv.Itoa(ctx.Writer.Status()),
	}).Observe(float64(time.Since(start).Seconds()))
}
