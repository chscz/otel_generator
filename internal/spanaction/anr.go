package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type AnrAttribute interface {
	ExceptionTypeRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	ExceptionMessageRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	ExceptionStackTraceRandomGenerate(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
}
type Anr struct {
	spanType attrspan.SpanAttrSpanType
	attr     AnrAttribute
}

func NewAnr(attrGenerator AnrAttribute) *Anr {
	return &Anr{spanType: attrspan.SpanAttrSpanTypeANR, attr: attrGenerator}
}

func (a *Anr) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		a.attr.ExceptionTypeRandomGenerate(a.spanType),
		a.attr.ExceptionMessageRandomGenerate(a.spanType),
		a.attr.ExceptionStackTraceRandomGenerate(a.spanType),
	}
	return attrs, fmt.Sprintf("anr!!")
}
