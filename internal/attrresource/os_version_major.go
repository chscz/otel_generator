package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	ResourceAttributeOSVersionMajorKey = attribute.Key("os.version_major")
)

func SetAttrOSVersionMajor(val string) attribute.KeyValue {
	return ResourceAttributeOSVersionMajorKey.String(val)
}
