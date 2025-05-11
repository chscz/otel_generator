package attrspan

import "go.opentelemetry.io/otel/attribute"

const SpanHTTPMethodKey = attribute.Key("http.method")

func HTTPMethodKey(val string) attribute.KeyValue {
	return SpanHTTPMethodKey.String(val)
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
