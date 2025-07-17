package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type WebVitalsAttribute interface{}
type WebVitals struct {
	attr WebVitalsAttribute
}

func NewWebVitals(attrGenerator WebVitalsAttribute) *WebVitals {
	return &WebVitals{attr: attrGenerator}
}

func (w *WebVitals) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{}
	return attrs, fmt.Sprintf("web vitals!!")
}
