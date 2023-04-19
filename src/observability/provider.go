package observability

import (
	"context"
	"log"

	"github.com/devopscorner/golang-adot/src/config"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
)

func InitObservability(router *gin.Engine) {
	InitOtel(router)
	InitMetrics(router)
}

func InitOtel(router *gin.Engine) {
	// Initialize the OTLP exporter
	ctx := context.Background()
	endpoint := config.OtelOtlpEndpoint()
	if endpoint == "" {
		endpoint = "0.0.0.0:4317"
	}

	exporter, err := otlptracegrpc.New(
		ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
	)
	if err != nil {
		log.Fatalf("Failed to create OTLP exporter: %v", err)
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
	err = exporter.Start(ctx)
	if err != nil {
		log.Fatalf("Failed to start OTLP exporter: %v", err)
	}
}

// GetTraceIDFromContext returns the trace ID from the provided context
func GetTraceIDFromContext(ctx context.Context) string {
	spanCtx := trace.SpanFromContext(ctx).SpanContext()
	if spanCtx.IsSampled() {
		return spanCtx.TraceID().String()
	}
	return ""
}
