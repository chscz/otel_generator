package generator

import (
	"context"
	"math/rand"

	attr2 "otel-generator/internal/attr"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("otel-generator")

type SpanGenerator struct {
	SpanTypes   []attr2.SpanAttrSpanType
	HTTPMethods []attr2.SpanAttrHTTPMethod
}

func NewSpanGenerator() *SpanGenerator {
	return &SpanGenerator{
		SpanTypes:   attr2.GenerateSpanTypeMocks(),
		HTTPMethods: attr2.GenerateHTTPMethodMocks(),
	}
}

func (s *SpanGenerator) GenerateSpan() *trace.Span {
	_, span := tracer.Start(context.Background(), "test-name")
	defer span.End()

	span.SetAttributes(
		attr2.SpanTypeKey(s.pickSpanTypeRandom()),
		attr2.HTTPMethodKey(s.pickHTTPMethodRandom()),
	)

	return &span
}

func (s *SpanGenerator) pickSpanTypeRandom() string {
	return string(s.SpanTypes[rand.Intn(len(s.SpanTypes))])
}

func (s *SpanGenerator) pickHTTPMethodRandom() string {
	return string(s.HTTPMethods[rand.Intn(len(s.HTTPMethods))])
}
