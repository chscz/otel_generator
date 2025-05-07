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
	tracer      trace.Tracer
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

func NewSpanGeneratorWithTracer(tracer trace.Tracer) *SpanGenerator {
	return &SpanGenerator{
		tracer:      tracer,                          // 주입된 Tracer 저장
		SpanTypes:   attr2.GenerateSpanTypeMocks(),   //
		HTTPMethods: attr2.GenerateHTTPMethodMocks(), //
	}
}

// 기존 GenerateSpan 메서드는 Span을 시작하고 바로 종료했습니다.
// Span의 생명주기를 호출자가 관리하도록 하거나,
// 여기서는 Span에 속성을 채우는 헬퍼 함수 형태로 변경하는 것을 제안합니다.

// 예시: 주어진 Span에 무작위 속성을 설정하는 메서드
func (s *SpanGenerator) PopulateSpanAttributes(span trace.Span) {
	span.SetAttributes(
		attr2.SpanTypeKey(s.pickSpanTypeRandom()),     //
		attr2.HTTPMethodKey(s.pickHTTPMethodRandom()), //
	)
}

// Span을 생성하고 바로 종료하는 간단한 예시 (필요하다면 사용)
// name: Span 이름, ctx: 부모 Context
func (s *SpanGenerator) CreateAndEndExampleSpan(ctx context.Context, name string) {
	// s.tracer는 NewSpanGeneratorWithTracer를 통해 주입된 tracer 사용
	_, span := s.tracer.Start(ctx, name)
	defer span.End() // 함수 종료 시 Span 종료

	s.PopulateSpanAttributes(span) // 위에서 정의한 속성 채우기 메서드 사용
}
