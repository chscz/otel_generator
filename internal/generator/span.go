package generator

import (
	"context"
	"math/rand"
	"time"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"
	"otel-generator/internal/config"
	"otel-generator/internal/spanaction"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type SpanGenerator struct {
	tracer                     trace.Tracer
	serviceType                attrresource.ServiceType
	attrGenerator              *attrspan.SpanAttrGenerator
	userID                     string
	actionGen                  *spanaction.ActionGenerator
	r                          *rand.Rand
	maxChildSpanCount          int
	maxSpanDurationMilliSecond int
}

func NewSpanGenerator(ctx context.Context, serviceType attrresource.ServiceType, cfg *config.Config, routineID int) *SpanGenerator {
	spanAttrGen := attrspan.NewSpanAttrGenerator(
		serviceType,
		cfg.SpanAttributes,
		cfg.UserCount,
		cfg.GenerateOption.MinSessionIDRefreshIntervalMinute,
		cfg.GenerateOption.MaxSessionIDRefreshIntervalMinute,
	)

	if cfg.GenerateOption.UseSessionIDRefresh {
		go spanAttrGen.StartSessionIDRefreshTimer(ctx)
	}

	return &SpanGenerator{
		serviceType:                spanAttrGen.ServiceType,
		attrGenerator:              spanAttrGen,
		userID:                     spanAttrGen.GetRandomUserID(),
		actionGen:                  spanaction.NewActionGenerator(spanAttrGen),
		r:                          rand.New(rand.NewSource(time.Now().UnixNano() + int64(routineID))),
		maxChildSpanCount:          cfg.GenerateOption.MaxChildSpanCount,
		maxSpanDurationMilliSecond: cfg.GenerateOption.MaxSpanDurationMilliSecond,
	}
}

func (s *SpanGenerator) GenerateTrace(mainCtx context.Context) {
	parentCtx, rootSpan, inheritedAttr := s.GenerateParentSpan(mainCtx)

	for i := 0; i < s.r.Intn(s.maxChildSpanCount); i++ {
		childSpan := s.GenerateChildSpan(parentCtx, inheritedAttr)
		randomDelay := time.Duration(s.r.Intn(s.maxSpanDurationMilliSecond)) * time.Millisecond
		time.Sleep(randomDelay)
		childSpan.End()
	}
	rootSpan.End()
}

func (s *SpanGenerator) GenerateParentSpan(parentCtx context.Context) (context.Context, trace.Span, attrspan.InheritedSpanAttr) {
	ctx, spanType, span := s.generateSpan(parentCtx)
	inheritedAttr := s.setPopulateParentSpanAttributes(span, spanType)
	return ctx, span, inheritedAttr
}

func (s *SpanGenerator) GenerateChildSpan(parentCtx context.Context, parentAttr attrspan.InheritedSpanAttr) trace.Span {
	_, spanType, span := s.generateSpan(parentCtx)
	s.setPopulateChildSpanAttributes(span, spanType, parentAttr)
	return span
}

func (s *SpanGenerator) generateSpan(parentCtx context.Context) (context.Context, attrspan.SpanAttrSpanType, trace.Span) {
	attrSpanType := attrspan.SpanAttrSpanType(s.attrGenerator.GenerateRandomSpanType().Value.AsString())
	attrs, spanName := s.actionGen.Generate(attrSpanType)

	var opt []trace.SpanStartOption
	if attrSpanType == attrspan.SpanAttrSpanTypeXHR {
		opt = append(opt, trace.WithSpanKind(trace.SpanKindClient))
	}

	ctx, span := s.tracer.Start(parentCtx, spanName, opt...)
	span.SetAttributes(attrs...)

	if attrSpanType.IsErrorSpanType() {
		span.SetStatus(codes.Error, "test-description!!")
	}

	return ctx, attrSpanType, span
}

func (s *SpanGenerator) setPopulateParentSpanAttributes(span trace.Span, spanType attrspan.SpanAttrSpanType) attrspan.InheritedSpanAttr {
	return s.attrGenerator.SetPopulateParentSpanAttributes(span, spanType, s.userID)
}

func (s *SpanGenerator) setPopulateChildSpanAttributes(span trace.Span, spanType attrspan.SpanAttrSpanType, attr attrspan.InheritedSpanAttr) {
	s.attrGenerator.SetPopulateChildSpanAttributes(span, spanType, attr)
}
