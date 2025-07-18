package attrspan

import (
	"math/rand"

	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

type SpanAttributeExceptionStackTrace struct {
	Crash []string `yaml:"crash"`
	ANR   []string `yaml:"anr"`
	Error []string `yaml:"error"`
}

func (es SpanAttributeExceptionStackTrace) GetAttributes(serviceType attrresource.ServiceType) []string {
	switch serviceType {
	case attrresource.ServiceTypeIOS:
		return append(es.Crash, es.Error...)
	case attrresource.ServiceTypeAndroid:
		return append(append(es.Crash, es.Error...), es.ANR...)
	case attrresource.ServiceTypeWeb:
		return es.Error
	default:
		return nil
	}
}

func (sg *SpanAttrGenerator) ExceptionStackTraceKey(val string) attribute.KeyValue {
	return semconv.ExceptionStacktrace(val)
}

func (sg *SpanAttrGenerator) ExceptionStackTraceRandomGenerate() attribute.KeyValue {
	return sg.ExceptionStackTraceKey(sg.ExceptionStackTraces[rand.Intn(len(sg.ExceptionStackTraces))])
}
