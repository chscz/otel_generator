package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func (sg *SpanAttrGenerator) ExceptionTypeKey(val string) attribute.KeyValue {
	return semconv.ExceptionType(val)
}

func (sg *SpanAttrGenerator) ExceptionTypeRandomGenerate() attribute.KeyValue {
	return sg.ExceptionTypeKey(sg.ExceptionTypes[rand.Intn(len(sg.ExceptionTypes))])
}
