package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"
)

func (sg *SpanAttrGenerator) ExceptionStackTraceKey(val string) attribute.KeyValue {
	return semconv.ExceptionStacktrace(val)
}

func (sg *SpanAttrGenerator) ExceptionStackTraceRandomGenerate() attribute.KeyValue {
	return sg.ExceptionStackTraceKey(sg.ExceptionStackTraces[rand.Intn(len(sg.ExceptionStackTraces))])
}
