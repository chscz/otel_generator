package attrspan

import (
	"math/rand"

	"otel-generator/internal/attrresource"

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

func (sg *SpanAttrGenerator) ExceptionMessageKey(val string) attribute.KeyValue {
	return semconv.ExceptionMessage(val)
}

func (sg *SpanAttrGenerator) ExceptionMessageRandomGenerate() attribute.KeyValue {
	return sg.ExceptionTypeKey(sg.ExceptionMessages[rand.Intn(len(sg.ExceptionMessages))])
}
