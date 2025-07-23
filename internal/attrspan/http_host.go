package attrspan

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/semconv/v1.12.0"
)

func (sg *SpanAttrGenerator) SetAttrHTTPHost(val string) attribute.KeyValue {
	return semconv.HTTPHostKey.String(val)
}
