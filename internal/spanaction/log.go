package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type LogAttribute interface{}
type Log struct {
	attr LogAttribute
}

func NewLog(attrGenerator LogAttribute) *Log {
	return &Log{attr: attrGenerator}
}

func (l *Log) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{}
	return attrs, fmt.Sprintf("log!!")
}
