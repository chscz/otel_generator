package generator

import (
	"context"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"
	"otel-generator/internal/config"
	"otel-generator/internal/spanaction"

	"go.opentelemetry.io/otel/trace"
)

type SpanGenerator struct {
	tracer        trace.Tracer
	platform      attrresource.PlatformType
	attrGenerator *attrspan.SpanAttrGenerator
	userID        string
	actionGen     *spanaction.ActionGenerator
}

func NewSpanGenerator(platform attrresource.PlatformType, cfg *config.Config) *SpanGenerator {
	spanAttrGen := attrspan.NewSpanAttrGenerator(cfg.SpanAttributes.ScreenNames, cfg.SpanAttributes.HTTPURLs, cfg.UserCount)

	return &SpanGenerator{
		platform:      platform,
		attrGenerator: spanAttrGen,
		userID:        spanAttrGen.GetRandomUserID(),
		actionGen:     spanaction.NewActionGenerator(spanAttrGen),
	}
}

func (s *SpanGenerator) GenerateTrace(mainCtx context.Context) {
	parentCtx, rootSpan := s.GenerateParentSpan(mainCtx)

	for i := 0; i < 5; i++ {
		childSpan := s.GenerateChildSpan(parentCtx)
		childSpan.End()
	}
	rootSpan.End()
}

func (s *SpanGenerator) GenerateParentSpan(parentCtx context.Context) (context.Context, trace.Span) {
	spanName := "root_span_name"
	attrSpanType := attrspan.SpanAttrSpanType(s.attrGenerator.SpanTypeRandomGenerate().Value.AsString())

	var taskCtx context.Context
	var rootSpan trace.Span
	if attrSpanType == attrspan.SpanAttrSpanTypeXHR {
		taskCtx, rootSpan = s.tracer.Start(parentCtx, spanName, trace.WithSpanKind(trace.SpanKindClient))
	} else {
		taskCtx, rootSpan = s.tracer.Start(parentCtx, spanName)
	}

	_ = s.populateSpanAttributes(rootSpan)

	switch attrSpanType {
	case attrspan.SpanAttrSpanTypeANR:
	case attrspan.SpanAttrSpanTypeCrash:
	case attrspan.SpanAttrSpanTypeError:
	case attrspan.SpanAttrSpanTypeEvent:
	case attrspan.SpanAttrSpanTypeLog:
	case attrspan.SpanAttrSpanTypeRender:
	case attrspan.SpanAttrSpanTypeXHR:
		s.actionGen.XHR.Generate(taskCtx, rootSpan)
	case attrspan.SpanAttrSpanTypeWebVitals:
	default:

	}

	return taskCtx, rootSpan

}

func (s *SpanGenerator) GenerateChildSpan(parentCtx context.Context) trace.Span {
	_, childSpan := s.tracer.Start(parentCtx, "child_span_name")
	s.populateSpanAttributes(childSpan)
	return childSpan
}

func (s *SpanGenerator) populateSpanAttributes(span trace.Span) attrspan.SpanAttrSpanType {
	return s.attrGenerator.PopulateSpanAttributes(span, s.userID)
}
