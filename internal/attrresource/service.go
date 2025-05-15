package attrresource

import "go.opentelemetry.io/otel/attribute"

type PlatformType string

const (
	PlatformTypeAndroid = "android"
	PlatformTypeIOS     = "ios"
	PlatformTypeWeb     = "web"
)

type Service struct {
	Name     string       `yaml:"name"`
	Version  string       `yaml:"version"`
	Key      string       `yaml:"key"`
	Platform PlatformType `yaml:"platform"`
}

const (
	ResourceAttributeServiceKey      = attribute.Key("service.key")
	ResourceAttributeServicePlatform = attribute.Key("service.platform")
)

func ServicePlatform(val PlatformType) attribute.KeyValue {
	return ResourceAttributeServicePlatform.String(string(val))
}

func ServiceKey(val string) attribute.KeyValue {
	return ResourceAttributeServiceKey.String(val)
}

func GenerateServiceMocks() []Service {
	return []Service{
		{Name: "test-service-ios-1", Version: "1.0.0", Key: "service-key", Platform: PlatformTypeIOS},
		{Name: "test-service-ios-1", Version: "1.0.1", Key: "service-key", Platform: PlatformTypeIOS},
		{Name: "test-service-ios-1", Version: "1.0.2", Key: "service-key", Platform: PlatformTypeIOS},
		{Name: "test-service-ios-2", Version: "1.0.11", Key: "service-key", Platform: PlatformTypeIOS},
		{Name: "test-service-android-1", Version: "1.3.5", Key: "service-key", Platform: PlatformTypeAndroid},
		{Name: "test-service-android-1", Version: "2.1.1", Key: "service-key", Platform: PlatformTypeAndroid},
		{Name: "test-service-android-2", Version: "2.1.1", Key: "service-key", Platform: PlatformTypeAndroid},
		{Name: "test-service-android-3", Version: "1.0.4", Key: "service-key", Platform: PlatformTypeAndroid},
		{Name: "test-service-web-1", Version: "3.0.11", Key: "service-key", Platform: PlatformTypeWeb},
		{Name: "test-service-web-1", Version: "3.0.12", Key: "service-key", Platform: PlatformTypeWeb},
		{Name: "test-service-web-2", Version: "3.2.1", Key: "service-key", Platform: PlatformTypeWeb},
	}
}
