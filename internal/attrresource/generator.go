package attrresource

type ResourceAttrGenerator struct {
	Services []Service
	OSNames  ResourceAttributeOSName
}

func NewResourceAttrGenerator(services []Service, resourceAttr ResourceAttributes) *ResourceAttrGenerator {
	return &ResourceAttrGenerator{
		Services: services,
		OSNames:  resourceAttr.OSNames,
	}
}
