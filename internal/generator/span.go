package generator

import (
	"math/rand"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer = otel.Tracer("otel-generator")

type SpanGenerator struct {
	tracer      trace.Tracer
	platform    attrresource.PlatformType
	SpanTypes   []attrspan.SpanAttrSpanType
	HTTPMethods []attrspan.SpanAttrHTTPMethod
	Platform    attrresource.PlatformType
}

func NewSpanGeneratorWithTracer(tracer trace.Tracer, platform attrresource.PlatformType) *SpanGenerator {
	return &SpanGenerator{
		tracer:      tracer,
		platform:    platform,
		SpanTypes:   attrspan.GenerateSpanTypeMocks(),
		HTTPMethods: attrspan.GenerateHTTPMethodMocks(),
	}
}

//func (s *SpanGenerator) GenerateSpan() *trace.Span {
//	_, span := tracer.Start(context.Background(), "test-name")
//	defer span.End()
//
//	span.SetAttributes(
//		attrspan.SpanTypeKey(s.pickSpanTypeRandom()),
//		attrspan.HTTPMethodKey(s.pickHTTPMethodRandom()),
//	)
//
//	return &span
//}

func (s *SpanGenerator) pickSpanTypeRandom() string {
	return string(s.SpanTypes[rand.Intn(len(s.SpanTypes))])
}

func (s *SpanGenerator) pickHTTPMethodRandom() string {
	return string(s.HTTPMethods[rand.Intn(len(s.HTTPMethods))])
}

func (s *SpanGenerator) PopulateSpanAttributes(span trace.Span) {
	span.SetAttributes(
		attrspan.SpanTypeKey(s.pickSpanTypeRandom()),
		attrspan.HTTPMethodKey(s.pickHTTPMethodRandom()),
	)
}

//func (s *SpanGenerator) CreateAndEndExampleSpan(ctx context.Context, name string) {
//	// s.tracer는 NewSpanGeneratorWithTracer를 통해 주입된 tracer 사용
//	_, span := s.tracer.Start(ctx, name)
//	defer span.End() // 함수 종료 시 Span 종료
//
//	s.PopulateSpanAttributes(span) // 위에서 정의한 속성 채우기 메서드 사용
//}
