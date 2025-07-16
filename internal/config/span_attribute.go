package config

import "otel-generator/internal/attrspan"

type SpanAttributes struct {
	ScreenNames          attrspan.SpanAttributeScreenName `yaml:"screen_names"`
	HTTPURLs             []string                         `yaml:"http_urls"`
	HTTPMethods          []string                         `yaml:"http_methods"`
	ExceptionTypes       []string                         `yaml:"exception_types"`
	ExceptionMessages    []string                         `yaml:"exception_messages"`
	ExceptionStackTraces []string                         `yaml:"exception_stack_traces"`
}
