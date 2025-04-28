package domain

type PlatformType string

const (
	PlatformTypeAndroid = "android"
	PlatformTypeIOS     = "ios"
	PlatformTypeWeb     = "web"
)

type Service struct {
	Name     string
	Version  string
	Key      string
	Platform PlatformType
}

func generateServiceMocks() map[int]Service {
	m := make(map[int]Service)
	m[0] = Service{Name: "test-service-ios-1", Version: "1.0.0", Key: "service-key", Platform: PlatformTypeIOS}
	m[1] = Service{Name: "test-service-ios-1", Version: "1.0.1", Key: "service-key", Platform: PlatformTypeIOS}
	m[2] = Service{Name: "test-service-ios-1", Version: "1.0.2", Key: "service-key", Platform: PlatformTypeIOS}
	m[3] = Service{Name: "test-service-ios-2", Version: "1.0.11", Key: "service-key", Platform: PlatformTypeIOS}
	m[4] = Service{Name: "test-service-android-1", Version: "1.3.5", Key: "service-key", Platform: PlatformTypeAndroid}
	m[5] = Service{Name: "test-service-android-1", Version: "2.1.1", Key: "service-key", Platform: PlatformTypeAndroid}
	m[6] = Service{Name: "test-service-android-2", Version: "2.1.1", Key: "service-key", Platform: PlatformTypeAndroid}
	m[7] = Service{Name: "test-service-android-3", Version: "1.0.4", Key: "service-key", Platform: PlatformTypeAndroid}
	m[8] = Service{Name: "test-service-web-1", Version: "3.0.11", Key: "service-key", Platform: PlatformTypeWeb}
	m[9] = Service{Name: "test-service-web-1", Version: "3.0.12", Key: "service-key", Platform: PlatformTypeWeb}
	m[10] = Service{Name: "test-service-web-2", Version: "3.2.1", Key: "service-key", Platform: PlatformTypeWeb}
	return m
}
