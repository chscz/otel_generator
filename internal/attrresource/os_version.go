package attrresource

import (
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

type ResourceAttributeOSVersion struct {
	Android []string `yaml:"android"`
	IOS     []string `yaml:"ios"`
	Web     []string `yaml:"web"`
}

func (ov ResourceAttributeOSVersion) GetAttributes(serviceType ServiceType) []string {
	switch serviceType {
	case ServiceTypeAndroid:
		return ov.Android
	case ServiceTypeIOS:
		return ov.IOS
	case ServiceTypeWeb:
		return ov.Web
	default:
		return nil
	}
}

func (rg *ResourceAttrGenerator) SetAttrOSVersion(val string) attribute.KeyValue {
	return semconv.OSVersion(val)
}

func (rg *ResourceAttrGenerator) GenerateRandomOSVersion(serviceType ServiceType) attribute.KeyValue {
	osVersions := GetAttributeByServiceType[ResourceAttributeOSVersion](serviceType, rg.OSVersions)
	osVersion, ok := util.PickRandomElementFromSlice[string](osVersions)
	if !ok {
		return attribute.KeyValue{}
	}
	return rg.SetAttrOSVersion(osVersion)
}
