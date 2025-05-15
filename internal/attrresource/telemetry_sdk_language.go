package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func TelemetrySDKLanguage(val string) attribute.KeyValue {
	switch val {
	case PlatformTypeAndroid:
		return semconv.TelemetrySDKLanguageJava
	case PlatformTypeIOS:
		return semconv.TelemetrySDKLanguageSwift
	case PlatformTypeWeb:
		return semconv.TelemetrySDKLanguageNodejs
	default:
		return semconv.TelemetrySDKLanguageKey.String("unknown")
	}
}
