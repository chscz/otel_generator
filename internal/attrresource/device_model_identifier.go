package attrresource

import (
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

type ResourceAttributeDeviceModelIdentifier struct {
	Android []string `yaml:"android"`
	IOS     []string `yaml:"ios"`
	Web     []string `yaml:"web"`
}

func (dm ResourceAttributeDeviceModelIdentifier) GetAttributes(serviceType ServiceType) []string {
	switch serviceType {
	case ServiceTypeAndroid:
		return dm.Android
	case ServiceTypeIOS:
		return dm.IOS
	case ServiceTypeWeb:
		return dm.Web
	default:
		return nil
	}
}

func (rg *ResourceAttrGenerator) SetAttrDeviceModelIdentifier(val string) attribute.KeyValue {
	return semconv.DeviceModelIdentifier(val)
}

func (rg *ResourceAttrGenerator) GenerateRandomDeviceModelIdentifier(serviceType ServiceType) attribute.KeyValue {
	deviceModelIdentifiers := GetAttributeByServiceType[ResourceAttributeDeviceModelIdentifier](serviceType, rg.DeviceModelIdentifiers)
	deviceModelIdentifier, ok := util.PickRandomElementFromSlice[string](deviceModelIdentifiers)
	if !ok {
		return attribute.KeyValue{}
	}
	return rg.SetAttrDeviceModelIdentifier(deviceModelIdentifier)

}
