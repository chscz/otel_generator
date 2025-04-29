package main

import (
	"context"
	"log"

	generator2 "otel-generator/internal/generator"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp" // OTLP HTTP 트레이스 익스포터
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	exporter, err := otlptracehttp.New(context.Background(),
		otlptracehttp.WithEndpointURL("http://localhost:4318/v1/trace"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("failed to create exporter: %v", err)
	}

	rg := generator2.NewResource(1000)
	sg := generator2.NewSpanGenerator()

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(rg.GenerateResource()),
	)

	otel.SetTracerProvider(tp)

	tg := generator2.NewTraceGenerator(rg, sg)
	for {
		resource, span := tg.GenerateTrace()
		_ = resource
		_ = span
	}
}
