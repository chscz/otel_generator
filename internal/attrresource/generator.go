package attrresource

import (
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
)

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
		Services:               services,
		OSNames:                resourceAttr.OSNames,
		OSVersions:             resourceAttr.OSVersions,
		DeviceModelIdentifiers: resourceAttr.DeviceModelIdentifier,
	}
}

func GetAttributeByServiceType[T AttributeSourceByServiceType](serviceType ServiceType, attr T) []string {
	return attr.GetAttributes(serviceType)
}

func (rg *ResourceAttrGenerator) SetPopulateAttribute(serviceType ServiceType) []attribute.KeyValue {
	osVersion := rg.GenerateRandomOSVersion(serviceType)
	major := util.GetMajorVersion(osVersion.Value.AsString())

	return []attribute.KeyValue{
		rg.GenerateRandomOSName(serviceType),
		osVersion,
		rg.SetAttrOSVersionMajor(major),
		rg.GenerateRandomDeviceModelIdentifier(serviceType),
	}
}
