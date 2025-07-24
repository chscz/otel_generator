package attrresource

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	ResourceAttributeOSVersionMajorKey = attribute.Key("os.version_major")
)

func (rg *ResourceAttrGenerator) SetAttrOSVersionMajor(val string) attribute.KeyValue {
	return ResourceAttributeOSVersionMajorKey.String(val)
}
