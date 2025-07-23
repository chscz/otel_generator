package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type AnrAttribute interface {
	GenerateRandomExceptionType(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	GenerateRandomExceptionMessage(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
	GenerateRandomExceptionStackTrace(spanType attrspan.SpanAttrSpanType) attribute.KeyValue
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
		a.attr.GenerateRandomExceptionType(a.spanType),
		a.attr.GenerateRandomExceptionMessage(a.spanType),
		a.attr.GenerateRandomExceptionStackTrace(a.spanType),
	}
	return attrs, fmt.Sprintf("anr!!")
}
