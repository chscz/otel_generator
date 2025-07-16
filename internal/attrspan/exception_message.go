package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func (sg *SpanAttrGenerator) ExceptionMessageKey(val string) attribute.KeyValue {
	return semconv.ExceptionMessage(val)
}

func (sg *SpanAttrGenerator) ExceptionMessageRandomGenerate() attribute.KeyValue {
	return sg.ExceptionTypeKey(sg.ExceptionMessages[rand.Intn(len(sg.ExceptionMessages))])
}
