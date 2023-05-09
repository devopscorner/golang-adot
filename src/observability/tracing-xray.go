package observability

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/aws/aws-xray-sdk-go/xraylog"
	"github.com/devopscorner/golang-adot/src/config"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type response struct {
	TraceID string `json:"traceId"`
}

func InitTracingXRay(router *gin.Engine) {
	// Initialize the XRay
	xray_endpoint := config.XRayDaemonEndpoint()
	if xray_endpoint == "" {
		xray_endpoint = "https://xray.us-west-2.amazonaws.com"
	}

	xray_version := config.XRayVersion()
	if xray_version == "" {
		xray_version = "latest"
	}

	xray.Configure(xray.Config{
		DaemonAddr:     xray_endpoint,
		ServiceVersion: xray_version,
		LogLevel:       config.LogLevel(),
	})

	// Configure X-Ray logging
	xray.SetLogger(xraylog.NewDefaultLogger(os.Stdout, xraylog.LogLevelInfo))

	router.Use(XrayMiddleware())
}

func XrayMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start X-Ray segment
		ctx, seg := xray.BeginSegment(c.Request.Context(), config.OtelServiceName())

		c.Request = c.Request.WithContext(ctx)
		// Continue with the request
		c.Next()

		// End X-Ray segment
		defer seg.Close(nil)
	}
}

// GetXrayTraceID generates a trace ID in Xray format from the span context.
func GetTraceIDXRay(span trace.Span) string {
	xrayTraceID := span.SpanContext().TraceID().String()
	return fmt.Sprintf("1-%s-%s", xrayTraceID[0:8], xrayTraceID[8:])
}

func WriteResponseXRay(span trace.Span, w http.ResponseWriter) {
	xrayTraceID := GetTraceIDXRay(span)
	payload, _ := json.Marshal(response{TraceID: xrayTraceID})
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func GetTraceIDXRaySegment(ctx context.Context, path string) (context.Context, *xray.Segment) {
	segmentCtx, segment := xray.BeginSubsegment(ctx, path)
	return segmentCtx, segment
}

func WriteResponseXRaySegment(segment *xray.Segment, w http.ResponseWriter) {
	xrayTraceID := segment.TraceID
	payload, _ := json.Marshal(response{TraceID: xrayTraceID})
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
