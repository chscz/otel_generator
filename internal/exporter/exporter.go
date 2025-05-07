package exporter

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

type Exporter struct {
}

// func NewExporter(mainCtx context.Context, collectorURL string) *otlptrace.Exporter {
func NewExporter(mainCtx context.Context, collectorURL string) *otlptrace.Exporter {
	exporter, err := otlptracehttp.New(mainCtx,
		otlptracehttp.WithEndpointURL(collectorURL),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Panicf("failed to create exporter: %v", err)
	}
	return exporter
}
