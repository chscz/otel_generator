package spanaction

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type XHRAttribute interface {
	HTTPMethodKey(val string) attribute.KeyValue
	HTTPMethodRandomGenerate() attribute.KeyValue
	HTTPURLKey(val string) attribute.KeyValue
	HTTPURLRandomGenerate() attribute.KeyValue
	HTTPStatusCodeKey(val int) attribute.KeyValue
	HTTPStatusCodeRandomGenerate() attribute.KeyValue
}

type XHR struct {
	Attr XHRAttribute
}

func NewXHR(attrGenerator XHRAttribute) *XHR {
	return &XHR{
		Attr: attrGenerator,
	}
}

func (x *XHR) SetSpanAttribute(span trace.Span) {
	span.SetAttributes(
		x.HTTPMethodRandomGenerate(),
		x.HTTPURLRandomGenerate(),
		x.HTTPStatusCodeRandomGenerate(),
	)
}

func (x *XHR) Generate(ctx context.Context, span trace.Span) {
	x.SetSpanAttribute(span)
}

func (x *XHR) HTTPMethodKey(val string) attribute.KeyValue {
	return x.Attr.HTTPMethodKey(val)
}

func (x *XHR) HTTPMethodRandomGenerate() attribute.KeyValue {
	return x.Attr.HTTPMethodRandomGenerate()
}

func (x *XHR) HTTPURLKey(val string) attribute.KeyValue {
	return x.Attr.HTTPURLKey(val)
}

func (x *XHR) HTTPURLRandomGenerate() attribute.KeyValue {
	return x.Attr.HTTPURLRandomGenerate()
}

func (x *XHR) HTTPStatusCodeKey(val int) attribute.KeyValue {
	return x.Attr.HTTPStatusCodeKey(val)
}

func (x *XHR) HTTPStatusCodeRandomGenerate() attribute.KeyValue {
	return x.Attr.HTTPStatusCodeRandomGenerate()
}
