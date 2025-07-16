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
	Attr CrashAttribute
}

func NewCrash(attrGenerator CrashAttribute) *Crash {
	return &Crash{Attr: attrGenerator}
}

func (c *Crash) Generate() ([]attribute.KeyValue, string) {
	attrs := []attribute.KeyValue{
		c.Attr.ExceptionTypeRandomGenerate(),
		c.Attr.ExceptionMessageRandomGenerate(),
		c.Attr.ExceptionStackTraceRandomGenerate(),
	}
	return attrs, fmt.Sprintf("crash!!")
}
