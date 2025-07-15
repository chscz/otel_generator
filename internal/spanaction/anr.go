package spanaction

import (
	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type AnrAttribute interface {
	SpanTypeKey(val attrspan.SpanAttrSpanType) attribute.KeyValue
}
type Anr struct {
	Attr AnrAttribute
}

func NewAnr(attrGenerator AnrAttribute) *Anr {
	return &Anr{
		Attr: attrGenerator,
	}
}

func (a *Anr) SpanTypeKey(val attrspan.SpanAttrSpanType) attribute.KeyValue {
	return a.Attr.SpanTypeKey(val)
}
