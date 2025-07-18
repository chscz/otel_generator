package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type EventAttribute interface{}

type Event struct {
	spanType attrspan.SpanAttrSpanType
	attr     EventAttribute
}

func NewEvent(attrGenerator EventAttribute) *Event {
	return &Event{spanType: attrspan.SpanAttrSpanTypeEvent, attr: attrGenerator}
}

func (e *Event) Generate() ([]attribute.KeyValue, string) {
	return nil, fmt.Sprintf("event!!")
}
