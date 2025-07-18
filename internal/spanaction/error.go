package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type ErrorAttribute interface {
	ExceptionTypeRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	ExceptionMessageRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	ExceptionStackTraceRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
}

type Error struct {
	spanType attrspan.SpanAttrSpanType
	attr     ErrorAttribute
}

func NewError(attrGenerator ErrorAttribute) *Error {
	return &Error{spanType: attrspan.SpanAttrSpanTypeError, attr: attrGenerator}
}

func (e *Error) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		e.attr.ExceptionTypeRandomGenerate(e.spanType),
		e.attr.ExceptionMessageRandomGenerate(e.spanType),
		e.attr.ExceptionStackTraceRandomGenerate(e.spanType),
	}
	return attrs, fmt.Sprintf("error!!")
}
