package generator

import (
	"fmt"
	"math/rand"

	"otel-generator/attr"
	"otel-generator/domain"

	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type Resource struct {
	Services   map[int]domain.Service
	SessionIDs []string
}

func (r *Resource) GenerateResource() *resource.Resource {
	service := r.pickServiceRandom()
	sessionID := r.pickSessionIDRandom()

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
		fmt.Println(err)
		return nil
	}
	return rs
}

func (r *Resource) pickService(n int) domain.Service {
	return r.Services[n]
}

func (r *Resource) pickServiceRandom() domain.Service {
	return r.Services[rand.Intn(len(r.Services))]
}

func (r *Resource) pickSessionIDRandom() string {
	return r.SessionIDs[rand.Intn(len(r.SessionIDs))]
}
