package attrresource

import (
	"strings"

	"otel-generator/internal/util"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
)

func OSNameKey(val string) attribute.KeyValue {
	return semconv.OSName(val)
}

type ResourceAttributeOSName struct {
	Android []string `yaml:"android"`
	IOS     []string `yaml:"ios"`
	Web     []string `yaml:"web"`
}

func (rg *ResourceAttrGenerator) SetOSNameAttr(serviceType ServiceType) attribute.KeyValue {
	var pick string
	var ok bool
	switch ServiceType(strings.ToUpper(string(serviceType))) {
	case ServiceTypeAndroid:
		pick, ok = util.RandomElementFromSlice(rg.OSNames.Android)
	case ServiceTypeIOS:
		pick, ok = util.RandomElementFromSlice(rg.OSNames.IOS)
	case ServiceTypeWeb:
		pick, ok = util.RandomElementFromSlice(rg.OSNames.Web)
	}
	if !ok {
		return attribute.KeyValue{}
	}
	return OSNameKey(pick)
}
