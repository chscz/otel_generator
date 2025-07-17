package spanaction

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
)

type CrashAttribute interface {
	ExceptionTypeRandomGenerate() attribute.KeyValue
	ExceptionMessageRandomGenerate() attribute.KeyValue
	ExceptionStackTraceRandomGenerate() attribute.KeyValue
}

type Crash struct {
	attr CrashAttribute
}

func NewCrash(attrGenerator CrashAttribute) *Crash {
	return &Crash{attr: attrGenerator}
}

func (c *Crash) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		c.attr.ExceptionTypeRandomGenerate(),
		c.attr.ExceptionMessageRandomGenerate(),
		c.attr.ExceptionStackTraceRandomGenerate(),
	}
	return attrs, fmt.Sprintf("crash!!")
}
