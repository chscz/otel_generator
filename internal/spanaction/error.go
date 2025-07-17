package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type ErrorAttribute interface {
	ExceptionTypeRandomGenerate() attribute.KeyValue
	ExceptionMessageRandomGenerate() attribute.KeyValue
	ExceptionStackTraceRandomGenerate() attribute.KeyValue
}

type Error struct {
	attr ErrorAttribute
}

func NewError(attrGenerator ErrorAttribute) *Error {
	return &Error{attr: attrGenerator}
}

func (e *Error) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		e.attr.ExceptionTypeRandomGenerate(),
		e.attr.ExceptionMessageRandomGenerate(),
		e.attr.ExceptionStackTraceRandomGenerate(),
	}
	return attrs, fmt.Sprintf("error!!")
}
