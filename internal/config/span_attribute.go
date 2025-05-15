package config

import "otel-generator/internal/attrspan"

type SpanAttributes struct {
	ScreenNames attrspan.SpanAttributeScreenName `yaml:"screen_names"`
	HTTPURLs    []string                         `yaml:"http_urls"`
}
