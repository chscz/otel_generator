package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type AnrAttribute interface {
	ExceptionTypeRandomGenerate() attribute.KeyValue
	ExceptionMessageRandomGenerate() attribute.KeyValue
	ExceptionStackTraceRandomGenerate() attribute.KeyValue
}
type Anr struct {
	attr AnrAttribute
}

func NewAnr(attrGenerator AnrAttribute) *Anr {
	return &Anr{attr: attrGenerator}
}

func (a *Anr) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		a.attr.ExceptionTypeRandomGenerate(),
		a.attr.ExceptionMessageRandomGenerate(),
		a.attr.ExceptionStackTraceRandomGenerate(),
	}
	return attrs, fmt.Sprintf("anr!!")
}
