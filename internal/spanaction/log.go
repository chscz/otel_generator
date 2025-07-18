package spanaction

import (
	"fmt"

	"otel-generator/internal/attrspan"

	"go.opentelemetry.io/otel/attribute"
)

type LogAttribute interface{}
type Log struct {
	spanType attrspan.SpanAttrSpanType
	attr     LogAttribute
}

func NewLog(attrGenerator LogAttribute) *Log {
	return &Log{spanType: attrspan.SpanAttrSpanTypeLog, attr: attrGenerator}
}

func (l *Log) Generate() ([]attribute.KeyValue, string) {
	return nil, fmt.Sprintf("log!!")
}
