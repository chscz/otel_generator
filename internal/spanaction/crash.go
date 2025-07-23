package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type CrashAttribute interface {
	GenerateRandomExceptionType(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	GenerateRandomExceptionMessage(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	GenerateRandomExceptionStackTrace(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
}

type Crash struct {
	spanType attrspan.SpanAttrSpanType
	attr     CrashAttribute
}

func NewCrash(attrGenerator CrashAttribute) *Crash {
	return &Crash{spanType: attrspan.SpanAttrSpanTypeCrash, attr: attrGenerator}
}

func (c *Crash) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		c.attr.GenerateRandomExceptionType(c.spanType),
		c.attr.GenerateRandomExceptionMessage(c.spanType),
		c.attr.GenerateRandomExceptionStackTrace(c.spanType),
	}
	return attrs, fmt.Sprintf("crash!!")
}
