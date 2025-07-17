package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type RenderAttribute interface{}
type Render struct {
	attr RenderAttribute
}

func NewRender(attrGenerator RenderAttribute) *Render {
	return &Render{attr: attrGenerator}
}

func (r *Render) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{}
	return attrs, fmt.Sprintf("web vitals!!")
}
