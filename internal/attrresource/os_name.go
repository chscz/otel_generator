package attrresource

import (
	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

type ResourceAttributeOSName struct {
	Android []string `yaml:"android"`
	IOS     []string `yaml:"ios"`
	Web     []string `yaml:"web"`
}

func (on ResourceAttributeOSName) GetAttributes(serviceType ServiceType) []string {
	switch serviceType {
	case ServiceTypeAndroid:
		return on.Android
	case ServiceTypeIOS:
		return on.IOS
	case ServiceTypeWeb:
		return on.Web
	default:
		return nil
	}
}

func (rg *ResourceAttrGenerator) SetAttrOSName(val string) attribute.KeyValue {
	return semconv.OSName(val)
}

func (rg *ResourceAttrGenerator) GenerateRandomOSName(serviceType ServiceType) attribute.KeyValue {
	osNames := GetAttributeByServiceType[ResourceAttributeOSName](serviceType, rg.OSNames)
	osName, ok := util.PickRandomElementFromSlice[string](osNames)
	if !ok {
		return attribute.KeyValue{}
	}
	return rg.SetAttrOSName(osName)
}
