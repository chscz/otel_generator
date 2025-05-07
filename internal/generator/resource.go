package generator

import (
	"fmt"
	"math/rand"

	attr "otel-generator/internal/attr"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type ResourceGenerator struct {
	Services   []attr.Service
	SessionIDs []string
}

func NewResource() *ResourceGenerator {
	return &ResourceGenerator{
		Services: attr.GenerateServiceMocks(),
		//SessionIDs: attr.GenerateSessionIDMocks(sessionCount),
	}
}

func (r *ResourceGenerator) GenerateResource() *resource.Resource {
	service := r.pickServiceRandom()
	sessionID := attr.GenerateSessionIDMocks()

	rs, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service.Name),
			semconv.ServiceVersion(service.Version),
			attr.ServiceKey(service.Key),
			attr.SessionIDKey(sessionID),
		),
	)
	if err != nil {
		fmt.Printf("failed to generate resource: %v", err)
		return nil
	}
	return rs
}

func (r *ResourceGenerator) pickService(n int) attr.Service {
	return r.Services[n]
}

func (r *ResourceGenerator) pickServiceRandom() attr.Service {
	return r.Services[rand.Intn(len(r.Services))]
}

func (r *ResourceGenerator) pickSessionIDRandom() string {
	return r.SessionIDs[rand.Intn(len(r.SessionIDs))]
}
