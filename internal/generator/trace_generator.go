package generator

import (
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/trace"
)

type TraceGenerator struct {
	Resource *ResourceGenerator
	Span     *SpanGenerator
}

func NewTraceGenerator(resource *ResourceGenerator, span *SpanGenerator) *TraceGenerator {
	return &TraceGenerator{
		Resource: resource,
		Span:     span,
	}
}

func (tg *TraceGenerator) GenerateTrace() (*resource.Resource, *trace.Span) {
	return tg.Resource.GenerateResource(), tg.Span.GenerateSpan()
}
