package attrspan

import (
	"go.opentelemetry.io/otel/attribute"
)

type InheritedSpanAttr struct {
	SessionID  attribute.KeyValue
	UserID     attribute.KeyValue
	SpanType   attribute.KeyValue
	ScreenName attribute.KeyValue
	ScreenType attribute.KeyValue
}

type SpanAttributes struct {
	ScreenNames          SpanAttributeScreenName          `yaml:"screen_names"`
	HTTPURLs             []string                         `yaml:"http_urls"`
	HTTPMethods          []string                         `yaml:"http_methods"`
	ExceptionTypes       SpanAttributeExceptionType       `yaml:"exception_types"`
	ExceptionMessages    SpanAttributeExceptionMessage    `yaml:"exception_messages"`
	ExceptionStackTraces SpanAttributeExceptionStackTrace `yaml:"exception_stack_traces"`
}
