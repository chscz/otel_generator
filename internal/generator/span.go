package generator

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"otel-generator/internal/attrresource"
	"otel-generator/internal/attrspan"
	"otel-generator/internal/config"
	"otel-generator/internal/spanaction"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

type SpanGenerator struct {
	tracer        trace.Tracer
	serviceType   attrresource.ServiceType
	attrGenerator *attrspan.SpanAttrGenerator
	userID        string
	actionGen     *spanaction.ActionGenerator
	r             *rand.Rand
}

func NewSpanGenerator(serviceType attrresource.ServiceType, cfg *config.Config) *SpanGenerator {
	spanAttrGen := attrspan.NewSpanAttrGenerator(
		serviceType,
		cfg.SpanAttributes.ScreenNames,
		cfg.SpanAttributes.HTTPURLs,
		cfg.SpanAttributes.ExceptionTypes,
		cfg.SpanAttributes.ExceptionMessages,
		cfg.SpanAttributes.ExceptionStackTraces,
		cfg.UserCount,
	)

	return &SpanGenerator{
		serviceType:   serviceType,
		attrGenerator: spanAttrGen,
		userID:        spanAttrGen.GetRandomUserID(),
		actionGen:     spanaction.NewActionGenerator(spanAttrGen),
		r:             rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s *SpanGenerator) GenerateTrace(mainCtx context.Context) {
	parentCtx, rootSpan, inheritedAttr := s.GenerateParentSpan(mainCtx)

	for i := 0; i < s.r.Intn(15); i++ {
		childSpan := s.GenerateChildSpan(parentCtx, inheritedAttr)
		randomDelay := time.Duration(s.r.Intn(321)) * time.Millisecond
		time.Sleep(randomDelay)
		childSpan.End()
	}
	rootSpan.End()
}

func (s *SpanGenerator) GenerateParentSpan(parentCtx context.Context) (context.Context, trace.Span, attrspan.InheritedSpanAttr) {
	attrSpanType := attrspan.SpanAttrSpanType(s.attrGenerator.SpanTypeRandomGenerate().Value.AsString())

	var attrs []attribute.KeyValue
	var spanName string
	switch attrSpanType {
	case attrspan.SpanAttrSpanTypeXHR:
		attrs, spanName = s.actionGen.XHR.Generate()
	case attrspan.SpanAttrSpanTypeCrash:
		attrs, spanName = s.actionGen.Crash.Generate()
	default:
		spanName = fmt.Sprintf("this is parent span: %s", attrSpanType)
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

	inheritedAttr := s.setPopulateParentSpanAttributes(rootSpan, attrSpanType)

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

func (s *SpanGenerator) GenerateChildSpan(parentCtx context.Context, attr attrspan.InheritedSpanAttr) trace.Span {
	_, childSpan := s.tracer.Start(parentCtx, "child_span_name")
	s.setPopulateChildSpanAttributes(childSpan, attr)
	return childSpan
}

func (s *SpanGenerator) setPopulateParentSpanAttributes(span trace.Span, spanType attrspan.SpanAttrSpanType) attrspan.InheritedSpanAttr {
	return s.attrGenerator.SetPopulateParentSpanAttributes(span, spanType, s.userID)
}

func (s *SpanGenerator) setPopulateChildSpanAttributes(span trace.Span, attr attrspan.InheritedSpanAttr) {
	s.attrGenerator.SetPopulateChildSpanAttributes(span, attr)
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
