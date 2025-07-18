package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type RenderAttribute interface{}
type Render struct {
	spanType attrspan.SpanAttrSpanType
	attr     RenderAttribute
}

func NewRender(attrGenerator RenderAttribute) *Render {
	return &Render{spanType: attrspan.SpanAttrSpanTypeRender, attr: attrGenerator}
}

func (r *Render) Generate() ([]attribute.KeyValue, string) {
	return nil, fmt.Sprintf("web vitals!!")
}
