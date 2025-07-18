package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type WebVitalsAttribute interface{}
type WebVitals struct {
	spanType attrspan.SpanAttrSpanType
	attr     WebVitalsAttribute
}

func NewWebVitals(attrGenerator WebVitalsAttribute) *WebVitals {
	return &WebVitals{spanType: attrspan.SpanAttrSpanTypeWebVitals, attr: attrGenerator}
}

func (w *WebVitals) Generate() ([]attribute.KeyValue, string) {
	return nil, fmt.Sprintf("web vitals!!")
}
