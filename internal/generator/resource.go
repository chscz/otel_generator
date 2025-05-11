package generator

import (
	"fmt"
	"math/rand"

	"otel-generator/internal/attrresource"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type ResourceGenerator struct {
	Services []attrresource.Service
}

func NewResource() *ResourceGenerator {
	return &ResourceGenerator{
		Services: attrresource.GenerateServiceMocks(),
	}
}

type ResourceInfo struct {
	ServiceName    string
	ServiceVersion string
	SessionID      string
	Platform       attrresource.PlatformType
}

func (r *ResourceGenerator) GenerateResource() (*resource.Resource, ResourceInfo) {
	service := r.pickServiceRandom()
	sessionID := attrresource.GenerateSessionIDMocks()

	rs, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service.Name),
			semconv.ServiceVersion(service.Version),
			attrresource.ServiceKey(service.Key),
			attrresource.ServicePlatform(service.Platform),
			attrresource.SessionIDKey(sessionID),
		),
	)
	if err != nil {
		fmt.Printf("failed to generate resource: %v", err)
		return nil, ResourceInfo{}
	}
	return rs, ResourceInfo{
		ServiceName:    service.Name,
		ServiceVersion: service.Version,
		SessionID:      sessionID,
		Platform:       service.Platform,
	}
}

func (r *ResourceGenerator) pickService(n int) attrresource.Service {
	return r.Services[n]
}

func (r *ResourceGenerator) pickServiceRandom() attrresource.Service {
	return r.Services[rand.Intn(len(r.Services))]
}

func (s ResourceInfo) String() string {
	return fmt.Sprintf("%s@%s::%s::%s", s.ServiceName, s.ServiceVersion, s.Platform, s.SessionID)
}
