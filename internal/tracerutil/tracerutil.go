package tracerutil

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var globalTracer trace.Tracer

func InitTracer(serviceName string) error {

	globalTracer = otel.Tracer(serviceName)

	return nil
}

func GetTracer() trace.Tracer {
	return globalTracer
}
