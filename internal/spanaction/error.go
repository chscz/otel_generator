package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type ErrorAttribute interface {
	GenerateRandomExceptionType(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	GenerateRandomExceptionMessage(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	GenerateRandomExceptionStackTrace(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
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
		e.attr.GenerateRandomExceptionType(e.spanType),
		e.attr.GenerateRandomExceptionMessage(e.spanType),
		e.attr.GenerateRandomExceptionStackTrace(e.spanType),
	}
	return attrs, fmt.Sprintf("error!!")
}
