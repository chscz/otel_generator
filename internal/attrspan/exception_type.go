package attrspan

import (
	"math/rand"

	"otel-generator/internal/attrresource"

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

func (sg *SpanAttrGenerator) ExceptionTypeRandomGenerate() attribute.KeyValue {
	return sg.ExceptionTypeKey(sg.ExceptionTypes[rand.Intn(len(sg.ExceptionTypes))])
}
