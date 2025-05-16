package spanaction

import (
	"otel-generator/internal/attrspan"
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
