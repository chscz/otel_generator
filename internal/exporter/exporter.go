package exporter

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

type Exporter struct {
	Exp *otlptrace.Exporter
}

func NewExporter(mainCtx context.Context, collectorURL string) *Exporter {
	exporter, err := otlptracehttp.New(mainCtx,
		otlptracehttp.WithEndpointURL(collectorURL),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Panicf("failed to create exporter: %v", err)
	}
	
	return &Exporter{Exp: exporter}
}

func (e *Exporter) Shutdown() {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Exp.Shutdown(shutdownCtx); err != nil {
		log.Printf("failed to shutdown exporter: %v", err)
	} else {
		log.Println("success shutdown exporter!!!")
	}
}
