package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func SetAttrTelemetrySDKName(val string) attribute.KeyValue {
	return semconv.TelemetrySDKName(val)
}
