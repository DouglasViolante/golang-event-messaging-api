package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var traceInstance trace.Tracer

func GetTracer() trace.Tracer {
	if traceInstance == nil {
		traceInstance = otel.Tracer("event-messaging-api")
	}
	return traceInstance
}