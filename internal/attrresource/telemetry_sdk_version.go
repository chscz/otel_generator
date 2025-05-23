package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func TelemetrySDKVersion(val string) attribute.KeyValue {
	switch val {
	default:
		return semconv.TelemetrySDKVersion(val)
	}
}
