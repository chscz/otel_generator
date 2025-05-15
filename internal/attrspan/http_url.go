package attrspan

import (
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

//type SpanAttributeHTTPURL struct {
//	HTTPURLs []string `mapstructure:"http_urls"`
//}

func (sg *SpanAttrGenerator) HTTPURLKey(val string) attribute.KeyValue {
	return semconv.URLFull(val)
}

func GenerateHTTPURLMocks() []string {
	return []string{
		"www.google.com",
		"www.github.com",
	}
}
