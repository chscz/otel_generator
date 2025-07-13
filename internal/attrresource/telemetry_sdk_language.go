package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func TelemetrySDKLanguage(val string) attribute.KeyValue {
	switch val {
	case ServiceTypeAndroid:
		return semconv.TelemetrySDKLanguageJava
	case ServiceTypeIOS:
		return semconv.TelemetrySDKLanguageSwift
	case ServiceTypeWeb:
		return semconv.TelemetrySDKLanguageNodejs
	default:
		return semconv.TelemetrySDKLanguageKey.String("unknown")
	}
}
