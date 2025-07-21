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
	attrs := []attribute.KeyValue{
		attribute.Key("target_element_text").String("foo_target_element_text!!"),
		attribute.Key("target_element_id").String("foo_target_element_id!!"),
		attribute.Key("target_element").String("foo_target_element!!"),
	}

	return attrs, fmt.Sprintf("Click")
}
