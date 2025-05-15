package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.25.0"
)

const SpanHTTPMethodKey = attribute.Key("http.method")

func (sg *SpanAttrGenerator) HTTPMethodKey(val string) attribute.KeyValue {
	return SpanHTTPMethodKey.String(val)
}

func (sg *SpanAttrGenerator) HTTPMethodRandomGenerate() attribute.KeyValue {
	methods := GenerateHTTPMethodMocks()
	method := methods[rand.Intn(len(methods))]
	return semconv.HTTPMethod(string(method))
}

type SpanAttrHTTPMethod string

const (
	SpanAttrHTTPMethodGET     = "GET"
	SpanAttrHTTPMethodPOST    = "POST"
	SpanAttrHTTPMethodPUT     = "PUT"
	SpanAttrHTTPMethodPATCH   = "PATCH"
	SpanAttrHTTPMethodDELETE  = "DELETE"
	SpanAttrHTTPMethodOPTIONS = "OPTIONS"
)

func GenerateHTTPMethodMocks() []SpanAttrHTTPMethod {
	return []SpanAttrHTTPMethod{
		SpanAttrHTTPMethodGET,
		SpanAttrHTTPMethodPOST,
		SpanAttrHTTPMethodPUT,
		SpanAttrHTTPMethodPATCH,
		SpanAttrHTTPMethodDELETE,
		SpanAttrHTTPMethodOPTIONS,
	}
}
