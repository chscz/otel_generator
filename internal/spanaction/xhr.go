package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type XHRAttribute interface {
	SetAttrHTTPMethod(val string) attribute.KeyValue
	GenerateRandomHTTPMethod() attribute.KeyValue
	SetAttrHTTPURL(url string) attribute.KeyValue
	GenerateRandomHTTPURL() []attribute.KeyValue
	SetAttrHTTPStatusCode(val int) attribute.KeyValue
	GenerateRandomHTTPStatusCode() attribute.KeyValue
}

type XHR struct {
	spanType attrspan.SpanAttrSpanType
	Attr     XHRAttribute
}

func NewXHR(attrGenerator XHRAttribute) *XHR {
	return &XHR{spanType: attrspan.SpanAttrSpanTypeXHR, Attr: attrGenerator}
}

func (x *XHR) Generate() ([]attribute.KeyValue, string) {
	method := x.Attr.GenerateRandomHTTPMethod()
	statusCode := x.Attr.GenerateRandomHTTPStatusCode()
	attrs := []attribute.KeyValue{
		method,
		statusCode,
	}
	attrs = append(attrs, x.Attr.GenerateRandomHTTPURL()...)
	return attrs, fmt.Sprintf("http %s", method.Value.AsString())
}
