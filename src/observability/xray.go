package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func StartTracing(ctx context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer(name)
	return tracer.Start(ctx, name)
}

func EndTracing(span trace.Span, code codes.Code, message string) {
	span.SetStatus(codes.Error, message)
	span.End()
}
