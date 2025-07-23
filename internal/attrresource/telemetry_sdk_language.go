package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func SetAttrTelemetrySDKLanguage(val string) attribute.KeyValue {
	switch ServiceType(val) {
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
