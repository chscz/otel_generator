package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type ActionGenerator struct {
	Anr       *Anr
	Crash     *Crash
	Error     *Error
	Event     *Event
	Log       *Log
	Render    *Render
	WebVitals *WebVitals
	XHR       *XHR
}

func NewActionGenerator(spanAttrGen *attrspan.SpanAttrGenerator) *ActionGenerator {
	return &ActionGenerator{
		Anr:       NewAnr(spanAttrGen),
		Crash:     NewCrash(spanAttrGen),
		Error:     NewError(spanAttrGen),
		Event:     NewEvent(spanAttrGen),
		Log:       NewLog(spanAttrGen),
		Render:    NewRender(spanAttrGen),
		WebVitals: NewWebVitals(spanAttrGen),
		XHR:       NewXHR(spanAttrGen),
	}
}

func (ag *ActionGenerator) Generate(spanType attrspan.SpanAttrSpanType) ([]attribute.KeyValue, string) {
	var attrs []attribute.KeyValue
	var spanName string
	switch spanType {
	case attrspan.SpanAttrSpanTypeXHR:
		attrs, spanName = ag.XHR.Generate()
	case attrspan.SpanAttrSpanTypeCrash:
		attrs, spanName = ag.Crash.Generate()
	default:
		spanName = fmt.Sprintf("this spanType is <%s>", spanType)
	}
	return attrs, spanName
}
