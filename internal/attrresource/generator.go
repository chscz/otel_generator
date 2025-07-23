package attrresource

import "go.opentelemetry.io/otel/attribute"

type AttributeSourceByServiceType interface {
	GetAttributes(serviceType ServiceType) []string
}
type ResourceAttrGenerator struct {
	Services               []Service
	OSNames                ResourceAttributeOSName
	OSVersions             ResourceAttributeOSVersion
	DeviceModelIdentifiers ResourceAttributeDeviceModelIdentifier
}

func NewResourceAttrGenerator(services []Service, resourceAttr ResourceAttributes) *ResourceAttrGenerator {
	return &ResourceAttrGenerator{
		Services: services,
		OSNames:  resourceAttr.OSNames,
	}
}

func GetAttributeByServiceType[T AttributeSourceByServiceType](serviceType ServiceType, attr T) []string {
	return attr.GetAttributes(serviceType)
}

func (rg *ResourceAttrGenerator) SetPopulateAttribute(serviceType ServiceType) []attribute.KeyValue {
	return []attribute.KeyValue{
		rg.GenerateRandomOSName(serviceType),
		rg.GenerateRandomOSVersion(serviceType),
		rg.GenerateRandomDeviceModelIdentifier(serviceType),
	}
}
