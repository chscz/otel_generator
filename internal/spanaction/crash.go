package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type CrashAttribute interface {
	ExceptionTypeRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	ExceptionMessageRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	ExceptionStackTraceRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
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
		c.attr.ExceptionTypeRandomGenerate(c.spanType),
		c.attr.ExceptionMessageRandomGenerate(c.spanType),
		c.attr.ExceptionStackTraceRandomGenerate(c.spanType),
	}
	return attrs, fmt.Sprintf("crash!!")
}
