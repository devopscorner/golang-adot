package observability

import (
	"context"
	"log"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

const ServiceName = "bookstore-adot"

var TestingId = ""

var Tracer = otel.Tracer("github.com/devopscorner/golang-adot/src/observability")

// Names for metric instruments
const TimeAlive = "time_alive"
const CpuUsage = "cpu_usage"
const TotalHeapSize = "total_heap_size"
const ThreadsActive = "threads_active"
const TotalBytesSent = "total_bytes_sent"
const TotalApiRequests = "total_api_requests"
const LatencyTime = "latency_time"

// Common attributes for traces and metrics (random, request)
var RequestMetricCommonLabels = []attribute.KeyValue{
	attribute.String("signal", "metric"),
	attribute.String("language", ServiceName),
	attribute.String("metricType", "request"),
}

var RandomMetricCommonLabels = []attribute.KeyValue{
	attribute.String("signal", "metric"),
	attribute.String("language", ServiceName),
	attribute.String("metricType", "random"),
}

var TraceCommonLabels = []attribute.KeyValue{
	attribute.String("signal", "trace"),
	attribute.String("language", ServiceName),
}

func InitTracingOtel(router *gin.Engine) {
	// Set up Tracing Otel
	router.Use(SetTracingOtel())
}

func InitOtel(router *gin.Engine) error {
	ctx := context.Background()

	// Initialize the OTLP exporter
	endpoint := config.OtelOtlpEndpoint()
	if endpoint == "" {
		endpoint = "0.0.0.0:4317"
	}

	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	if err != nil {
		return err
	}

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exporter),
		sdktrace.WithResource(resource.NewSchemaless(semconv.ServiceNameKey.String(config.OtelServiceName()))),
	)

	// Set the global tracer
	otel.SetTracerProvider(tracerProvider)

	// Register the tracer with the Gin framework
	router.Use(otelgin.Middleware(config.OtelServiceName(), otelgin.WithTracerProvider(tracerProvider)))

	// Start the exporter
	if err := exporter.Start(ctx); err != nil {
		if err.Error() == "already started" {
			log.Println("OTLP exporter is already started")
		} else {
			return err
		}
	}

	log.Println("OTLP exporter is started")

	return nil
}

func SetTracingOtel() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Start tracing
		_, span := StartTracingOtel(ctx.Request.Context(), config.OtelServiceName())
		defer EndTracingOtel(span, codes.Ok, config.OtelServiceName())

		CalcRequests(ctx)
	}
}

func StartTracingOtel(ctx context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer(name)
	return tracer.Start(ctx, name)
}

func EndTracingOtel(span trace.Span, code codes.Code, message string) {
	span.SetStatus(codes.Error, message)
	span.End()
}

// GetTraceIDFromContext returns the trace ID from the provided context
func GetTraceIDFromContext(ctx context.Context) string {
	spanCtx := trace.SpanFromContext(ctx).SpanContext()
	if spanCtx.IsSampled() {
		return spanCtx.TraceID().String()
	}
	return ""
}
