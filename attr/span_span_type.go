package attr

import "go.opentelemetry.io/otel/attribute"

const (
	SpanSpanTypeKey = attribute.Key("span.type")
)

func SpanTypeKey(val string) attribute.KeyValue {
	return SpanSpanTypeKey.String(val)
}
