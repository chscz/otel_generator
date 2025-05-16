package spanaction

import "go.opentelemetry.io/otel/attribute"

type AnrAttribute interface {
	SpanTypeKey(val string) attribute.KeyValue
}
type Anr struct {
	Attr AnrAttribute
}

func NewAnr(attrGenerator AnrAttribute) *Anr {
	return &Anr{
		Attr: attrGenerator,
	}
}

func (a *Anr) SpanTypeKey(val string) attribute.KeyValue {
	return a.Attr.SpanTypeKey(val)
}
