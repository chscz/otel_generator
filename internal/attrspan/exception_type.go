package attrspan

import (
	"strings"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

type SpanAttributeExceptionType struct {
	Crash []string `yaml:"crash"`
	ANR   []string `yaml:"anr"`
	Error []string `yaml:"error"`
}

func (et SpanAttributeExceptionType) GetAttributes(serviceType attrresource.ServiceType) []string {
	switch serviceType {
	case attrresource.ServiceTypeIOS:
		return append(et.Crash, et.Error...)
	case attrresource.ServiceTypeAndroid:
		return append(append(et.Crash, et.Error...), et.ANR...)
	case attrresource.ServiceTypeWeb:
		return et.Error
	default:
		return nil
	}
}

func (sg *SpanAttrGenerator) ExceptionTypeKey(val string) attribute.KeyValue {
	return semconv.ExceptionType(val)
}

func (sg *SpanAttrGenerator) ExceptionTypeRandomGenerate(spanType SpanAttrSpanType) attribute.KeyValue {
	var exceptionTypesBySpanType []string
	for _, et := range sg.ExceptionTypes {
		if strings.HasPrefix(et, string(spanType)) {
			exceptionTypesBySpanType = append(exceptionTypesBySpanType, et)
		}
	}
	pick, ok := util.RandomElementFromSlice[string](exceptionTypesBySpanType)
	if !ok {
		return attribute.KeyValue{}
	}
	return sg.ExceptionTypeKey(pick)
}
