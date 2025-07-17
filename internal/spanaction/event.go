package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type EventAttribute interface{}

type Event struct {
	attr EventAttribute
}

func NewEvent(attrGenerator EventAttribute) *Event {
	return &Event{attr: attrGenerator}
}

func (e *Event) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{}
	return attrs, fmt.Sprintf("event!!")
}
