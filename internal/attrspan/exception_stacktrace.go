package attrspan

import (
	"strings"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/util"

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

func (sg *SpanAttrGenerator) SetAttrExceptionStackTrace(val string) attribute.KeyValue {
	return semconv.ExceptionStacktrace(val)
}

func (sg *SpanAttrGenerator) GenerateRandomExceptionStackTrace(spanType SpanAttrSpanType) attribute.KeyValue {
	var exceptionStackTraceBySpanType []string
	for _, et := range sg.ExceptionStackTraces {
		if strings.HasPrefix(et, string(spanType)) {
			exceptionStackTraceBySpanType = append(exceptionStackTraceBySpanType, et)
		}
	}
	pick, ok := util.PickRandomElementFromSlice[string](exceptionStackTraceBySpanType)
	if !ok {
		return attribute.KeyValue{}
	}
	return sg.SetAttrExceptionStackTrace(pick)
}
