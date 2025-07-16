package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type XHRAttribute interface {
	HTTPMethodKey(val string) attribute.KeyValue
	HTTPMethodRandomGenerate() attribute.KeyValue
	HTTPURLKey(url, host, method string) []attribute.KeyValue
	HTTPURLRandomGenerate() []attribute.KeyValue
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

//func (x *XHR) SetSpanAttribute(span trace.Span) {
//	span.SetAttributes(x.HTTPURLRandomGenerate()...)
//}

func (x *XHR) Generate() ([]attribute.KeyValue, string) {
	method := x.Attr.HTTPMethodRandomGenerate()
	attrs := []attribute.KeyValue{
		method,
		x.Attr.HTTPStatusCodeRandomGenerate(),
	}
	attrs = append(attrs, x.Attr.HTTPURLRandomGenerate()...)
	return attrs, fmt.Sprintf("http %s", method.Value.AsString())
}

//func (x *XHR) HTTPMethodKey(val string) attribute.KeyValue {
//	return x.Attr.HTTPMethodKey(val)
//}

//func (x *XHR) HTTPMethodRandomGenerate() attribute.KeyValue {
//	return x.Attr.HTTPMethodRandomGenerate()
//}

//func (x *XHR) HTTPURLKey(url, host, method string) []attribute.KeyValue {
//	return x.Attr.HTTPURLKey(url, host, method)
//}

//func (x *XHR) HTTPURLRandomGenerate() []attribute.KeyValue {
//	return x.Attr.HTTPURLRandomGenerate()
//}

//func (x *XHR) HTTPStatusCodeKey(val int) attribute.KeyValue {
//	return x.Attr.HTTPStatusCodeKey(val)
//}

//func (x *XHR) HTTPStatusCodeRandomGenerate() attribute.KeyValue {
//	return x.Attr.HTTPStatusCodeRandomGenerate()
//}
