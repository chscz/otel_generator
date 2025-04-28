package attr

import "go.opentelemetry.io/otel/attribute"

const (
	ResourceSessionIDKey = attribute.Key("session.id")
)

func SessionIDKey(val string) attribute.KeyValue {
	return ResourceSessionIDKey.String(val)
}
