package generator

import (
	"fmt"
	"strings"

	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type ResourceGenerator struct {
	Services      []attrresource.Service
	attrGenerator *attrresource.ResourceAttrGenerator
}

func NewResourceGenerator(services []attrresource.Service, attr attrresource.ResourceAttributes) *ResourceGenerator {
	resourceAttrGen := attrresource.NewResourceAttrGenerator(services, attr)
	return &ResourceGenerator{
		Services:      services,
		attrGenerator: resourceAttrGen,
	}
}

type ServiceInfo struct {
	ServiceName    string
	ServiceVersion string
	ServiceType    attrresource.ServiceType
}

func (r *ResourceGenerator) GenerateResource() (*resource.Resource, ServiceInfo) {
	service := r.attrGenerator.PickServiceRandom()

	attrs := []attribute.KeyValue{
		semconv.ServiceName(service.Name),
		semconv.ServiceVersion(service.Version),
		r.attrGenerator.SetAttrServiceKey(service.Key),
		r.attrGenerator.SetAttrServiceType(service.Type),
	}

	populateAttrs := r.attrGenerator.SetPopulateAttribute(service.Type)
	attrs = append(attrs, populateAttrs...)

	rs, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			attrs...,
		),
	)
	if err != nil {
		fmt.Printf("failed to generate resource: %v", err)
		return nil, ServiceInfo{}
	}
	return rs, ServiceInfo{
		ServiceName:    service.Name,
		ServiceVersion: service.Version,
		ServiceType:    attrresource.ServiceType(strings.ToUpper(string(service.Type))),
	}
}

func (s ServiceInfo) String() string {
	return fmt.Sprintf("%s@%s::%s", s.ServiceName, s.ServiceVersion, s.ServiceType)
}
