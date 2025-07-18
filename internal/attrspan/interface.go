package attrspan

import "otel-generator/internal/attrresource"

type AttributeSourceByServiceType interface {
	GetAttributes(serviceType attrresource.ServiceType) []string
}
