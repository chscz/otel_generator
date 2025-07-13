package generator

import (
	"context"
	"math/rand"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"
	"otel-generator/internal/config"
	"otel-generator/internal/spanaction"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type InheritedSpanAttr struct {
	//SessionID  attribute.KeyValue
	UserID     attribute.KeyValue
	SpanType   attribute.KeyValue
	ScreenName attribute.KeyValue
	ScreenType attribute.KeyValue
}

type SpanGenerator struct {
	tracer        trace.Tracer
	serviceType   attrresource.ServiceType
	attrGenerator *attrspan.SpanAttrGenerator
	userID        string
	actionGen     *spanaction.ActionGenerator
}

func NewSpanGenerator(platform attrresource.ServiceType, cfg *config.Config) *SpanGenerator {
	spanAttrGen := attrspan.NewSpanAttrGenerator(
		platform,
		cfg.SpanAttributes.ScreenNames,
		cfg.SpanAttributes.HTTPURLs,
		cfg.SpanAttributes.HTTPMethods,
		cfg.UserCount,
	)

	return &SpanGenerator{
		serviceType:   platform,
		attrGenerator: spanAttrGen,
		userID:        spanAttrGen.GetRandomUserID(),
		actionGen:     spanaction.NewActionGenerator(spanAttrGen),
	}
}

func (s *SpanGenerator) GenerateTrace(mainCtx context.Context) {
	parentCtx, rootSpan, inheritedAttr := s.GenerateParentSpan(mainCtx)

	for i := 0; i < rand.Intn(10); i++ {
		childSpan := s.GenerateChildSpan(parentCtx, inheritedAttr)
		childSpan.End()
	}
	rootSpan.End()
}

func (s *SpanGenerator) GenerateParentSpan(parentCtx context.Context) (context.Context, trace.Span, InheritedSpanAttr) {
	attrSpanType := attrspan.SpanAttrSpanType(s.attrGenerator.SpanTypeRandomGenerate().Value.AsString())

	var attrs []attribute.KeyValue
	var spanName string
	switch attrSpanType {
	case attrspan.SpanAttrSpanTypeXHR:
		attrs, spanName = s.actionGen.XHR.Generate()
	}

	var taskCtx context.Context
	var rootSpan trace.Span
	if attrSpanType == attrspan.SpanAttrSpanTypeXHR {
		taskCtx, rootSpan = s.tracer.Start(parentCtx, spanName, trace.WithSpanKind(trace.SpanKindClient))
		rootSpan.SetAttributes(attrs...)
	} else if (attrSpanType == attrspan.SpanAttrSpanTypeCrash) ||
		(attrSpanType == attrspan.SpanAttrSpanTypeANR) ||
		(attrSpanType == attrspan.SpanAttrSpanTypeError) {
		taskCtx, rootSpan = s.tracer.Start(parentCtx, spanName)
		rootSpan.SetAttributes(attrs...)
		rootSpan.SetStatus(codes.Error, "test-description!!")
	} else {
		taskCtx, rootSpan = s.tracer.Start(parentCtx, spanName)
		rootSpan.SetAttributes(attrs...)
	}

	inheritedAttr := s.setPopulateSpanAttributes(rootSpan, attrSpanType)

	//switch attrSpanType {
	//case attrspan.SpanAttrSpanTypeANR:
	//case attrspan.SpanAttrSpanTypeCrash:
	//case attrspan.SpanAttrSpanTypeError:
	//case attrspan.SpanAttrSpanTypeEvent:
	//case attrspan.SpanAttrSpanTypeLog:
	//case attrspan.SpanAttrSpanTypeRender:
	//case attrspan.SpanAttrSpanTypeXHR:
	//	//s.actionGen.XHR.Generate(taskCtx, rootSpan)
	//case attrspan.SpanAttrSpanTypeWebVitals:
	//default:
	//
	//}

	return taskCtx, rootSpan, inheritedAttr

}

func (s *SpanGenerator) GenerateChildSpan(parentCtx context.Context, attr InheritedSpanAttr) trace.Span {
	_, childSpan := s.tracer.Start(parentCtx, "child_span_name")
	s.setPopulateSpanAttributes(childSpan)
	return childSpan
}

func (s *SpanGenerator) setPopulateSpanAttributes(span trace.Span, spanType attrspan.SpanAttrSpanType) InheritedSpanAttr {
	return s.attrGenerator.SetPopulateSpanAttributes(span, spanType, s.userID)
}

func makeSpanNameBySpanType(spanType attrspan.SpanAttrSpanType) string {
	switch spanType {
	case attrspan.SpanAttrSpanTypeXHR:
		return "this is xhr"
	case attrspan.SpanAttrSpanTypeRender:
		return "this is render span"
	case attrspan.SpanAttrSpanTypeEvent:
		return "this is event span"
	case attrspan.SpanAttrSpanTypeCrash:
		return "this is crash span"
	case attrspan.SpanAttrSpanTypeError:
		return "this is error span"
	case attrspan.SpanAttrSpanTypeANR:
		return "this is anr span"
	case attrspan.SpanAttrSpanTypeLog:
		return "this is log span"
	case attrspan.SpanAttrSpanTypeWebVitals:
		return "this is web vital span"
	default:
		return "this is unknown span"
	}
}
