package attrresource

type ResourceAttributes struct {
	OSNames               ResourceAttributeOSName                `yaml:"os_name"`
	OSVersions            ResourceAttributeOSVersion             `yaml:"os_version"`
	DeviceModelIdentifier ResourceAttributeDeviceModelIdentifier `yaml:"device_model_identifier"`
}
