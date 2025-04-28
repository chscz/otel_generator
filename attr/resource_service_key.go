package attr

import "go.opentelemetry.io/otel/attribute"

const (
	ResourceServiceKey = attribute.Key("service.key")
)

func ServiceKey(val string) attribute.KeyValue {
	return ResourceServiceKey.String(val)
}
