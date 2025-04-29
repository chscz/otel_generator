package generator

import (
	"fmt"
	"math/rand"

	attr2 "otel-generator/internal/attr"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type ResourceGenerator struct {
	Services   []attr2.Service
	SessionIDs []string
}

func NewResource(sessionCount int) *ResourceGenerator {
	return &ResourceGenerator{
		Services:   attr2.GenerateServiceMocks(),
		SessionIDs: attr2.GenerateSessionIDMocks(sessionCount),
	}
}

func (r *ResourceGenerator) GenerateResource() *resource.Resource {
	service := r.pickServiceRandom()
	sessionID := r.pickSessionIDRandom()

	rs, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(service.Name),
			semconv.ServiceVersion(service.Version),
			attr2.ServiceKey(service.Key),
			attr2.SessionIDKey(sessionID),
		),
	)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rs
}

func (r *ResourceGenerator) pickService(n int) attr2.Service {
	return r.Services[n]
}

func (r *ResourceGenerator) pickServiceRandom() attr2.Service {
	return r.Services[rand.Intn(len(r.Services))]
}

func (r *ResourceGenerator) pickSessionIDRandom() string {
	return r.SessionIDs[rand.Intn(len(r.SessionIDs))]
}
