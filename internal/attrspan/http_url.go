package attrspan

import (
	"math/rand"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

//type SpanAttributeHTTPURL struct {
//	HTTPURLs []string `mapstructure:"http_urls"`
//}

func (sg *SpanAttrGenerator) HTTPURLKey(val string) attribute.KeyValue {
	return semconv.URLFull(val)
}

func (sg *SpanAttrGenerator) HTTPURLRandomGenerate() attribute.KeyValue {
	return semconv.URLFull(sg.HTTPURL[rand.Intn(len(sg.HTTPURL))])
}
