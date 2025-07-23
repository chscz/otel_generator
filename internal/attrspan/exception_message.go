package attrspan

import (
	"strings"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

type SpanAttributeExceptionMessage struct {
	Crash []string `yaml:"crash"`
	ANR   []string `yaml:"anr"`
	Error []string `yaml:"error"`
}

func (em SpanAttributeExceptionMessage) GetAttributes(serviceType attrresource.ServiceType) []string {
	switch serviceType {
	case attrresource.ServiceTypeIOS:
		return append(em.Crash, em.Error...)
	case attrresource.ServiceTypeAndroid:
		return append(append(em.Crash, em.Error...), em.ANR...)
	case attrresource.ServiceTypeWeb:
		return em.Error
	default:
		return nil
	}
}

func (sg *SpanAttrGenerator) SetAttrExceptionMessage(val string) attribute.KeyValue {
	return semconv.ExceptionMessage(val)
}

func (sg *SpanAttrGenerator) GenerateRandomExceptionMessage(spanType SpanAttrSpanType) attribute.KeyValue {
	var exceptionMessagesBySpanType []string
	for _, et := range sg.ExceptionMessages {
		if strings.HasPrefix(et, string(spanType)) {
			exceptionMessagesBySpanType = append(exceptionMessagesBySpanType, et)
		}
	}
	pick, ok := util.PickRandomElementFromSlice[string](exceptionMessagesBySpanType)
	if !ok {
		return attribute.KeyValue{}
	}
	return sg.SetAttrExceptionMessage(pick)
}
