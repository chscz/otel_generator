package attrspan

import (
	"net/url"

	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv12 "go.opentelemetry.io/otel/semconv/v1.12.0"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func (sg *SpanAttrGenerator) SetAttrHTTPURL(val string) attribute.KeyValue {
	return semconv.URLFull(val)
}

func (sg *SpanAttrGenerator) GenerateRandomHTTPURL() []attribute.KeyValue {
	urlFull, ok := util.PickRandomElementFromSlice[string](sg.HTTPURLs)
	if !ok {
		return nil
	}

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
