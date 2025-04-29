package generator

import (
	"context"
	"math/rand"

	"otel-generator/attr"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("otel-generator")

type SpanGenerator struct {
	SpanTypes   []attr.SpanAttrSpanType
	HTTPMethods []attr.SpanAttrHTTPMethod
}

func NewSpanGenerator() *SpanGenerator {
	return &SpanGenerator{
		SpanTypes:   attr.GenerateSpanTypeMocks(),
		HTTPMethods: attr.GenerateHTTPMethodMocks(),
	}
}

func (s *SpanGenerator) GenerateSpan() *trace.Span {
	_, span := tracer.Start(context.Background(), "test-name")
	defer span.End()

	span.SetAttributes(
		attr.SpanTypeKey(s.pickSpanTypeRandom()),
		attr.HTTPMethodKey(s.pickHTTPMethodRandom()),
	)

	return &span
}

func (s *SpanGenerator) pickSpanTypeRandom() string {
	return string(s.SpanTypes[rand.Intn(len(s.SpanTypes))])
}

func (s *SpanGenerator) pickHTTPMethodRandom() string {
	return string(s.HTTPMethods[rand.Intn(len(s.HTTPMethods))])
}
