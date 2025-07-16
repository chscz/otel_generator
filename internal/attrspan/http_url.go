package attrspan

import (
	"math/rand"
	"net/url"

	"go.opentelemetry.io/otel/attribute"
	semconv12 "go.opentelemetry.io/otel/semconv/v1.12.0"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

//type SpanAttributeHTTPURL struct {
//	HTTPURLs []string `mapstructure:"http_urls"`
//}

func (sg *SpanAttrGenerator) HTTPURLKey(url, host, method string) []attribute.KeyValue {

	return []attribute.KeyValue{
		semconv.URLFull(url),
		semconv12.HTTPHostKey.String(host),
		semconv12.HTTPMethodKey.String(method),
	}
}

func (sg *SpanAttrGenerator) HTTPURLRandomGenerate() []attribute.KeyValue {
	urlFull := sg.HTTPURLs[rand.Intn(len(sg.HTTPURLs))]

	parsedURL, err := url.Parse(urlFull)
	host := ""
	if err == nil && parsedURL.Host != "" {
		host = parsedURL.Host
	} else {
		tempURL, err2 := url.Parse("http://" + urlFull)
		if err2 == nil {
			host = tempURL.Host
		}
	}

	attrs := []attribute.KeyValue{
		semconv.URLFull(urlFull),
	}
	if host != "" {
		attrs = append(attrs, semconv12.HTTPHostKey.String(host))
	}

	return attrs
}
