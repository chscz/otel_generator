package attrresource

import (
	"math/rand"
	"strings"

	"go.opentelemetry.io/otel/attribute"
)

type ServiceType string

const (
	ServiceTypeAndroid ServiceType = "ANDROID"
	ServiceTypeIOS     ServiceType = "IOS"
	ServiceTypeWeb     ServiceType = "WEB"
)

type Service struct {
	Name    string      `yaml:"name"`
	Version string      `yaml:"version"`
	Type    ServiceType `yaml:"type"`
	Key     string      `yaml:"key"`
}

const (
	ResourceAttributeServiceKey = attribute.Key("service.key")
	//ResourceAttributeServiceType = attribute.Key("service.platform")
	ResourceAttributeServiceType = attribute.Key("imqa.service.type")
)

func (rg *ResourceAttrGenerator) SetAttrServiceType(val ServiceType) attribute.KeyValue {
	return ResourceAttributeServiceType.String(string(val))
}

func (rg *ResourceAttrGenerator) SetAttrServiceKey(val string) attribute.KeyValue {
	return ResourceAttributeServiceKey.String(val)
}

func GenerateServiceMocks() []Service {
	return []Service{
		{Name: "test-service-ios-1", Version: "1.0.0", Key: "service-key", Type: ServiceTypeIOS},
		{Name: "test-service-ios-1", Version: "1.0.1", Key: "service-key", Type: ServiceTypeIOS},
		{Name: "test-service-ios-1", Version: "1.0.2", Key: "service-key", Type: ServiceTypeIOS},
		{Name: "test-service-ios-2", Version: "1.0.11", Key: "service-key", Type: ServiceTypeIOS},
		{Name: "test-service-android-1", Version: "1.3.5", Key: "service-key", Type: ServiceTypeAndroid},
		{Name: "test-service-android-1", Version: "2.1.1", Key: "service-key", Type: ServiceTypeAndroid},
		{Name: "test-service-android-2", Version: "2.1.1", Key: "service-key", Type: ServiceTypeAndroid},
		{Name: "test-service-android-3", Version: "1.0.4", Key: "service-key", Type: ServiceTypeAndroid},
		{Name: "test-service-web-1", Version: "3.0.11", Key: "service-key", Type: ServiceTypeWeb},
		{Name: "test-service-web-1", Version: "3.0.12", Key: "service-key", Type: ServiceTypeWeb},
		{Name: "test-service-web-2", Version: "3.2.1", Key: "service-key", Type: ServiceTypeWeb},
	}
}

func (rg *ResourceAttrGenerator) PickServiceRandom() Service {
	service := rg.Services[rand.Intn(len(rg.Services))]
	service.Type = ServiceType(strings.ToUpper(string(service.Type)))
	return service
}
