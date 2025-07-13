package generator

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"otel-generator/internal/config"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const traceIntervalSeconds = 10

type TraceGenerator struct {
	routineID   int
	exporter    *otlptrace.Exporter
	resGen      *ResourceGenerator
	spanGen     *SpanGenerator
	tp          *sdktrace.TracerProvider
	serviceInfo ResourceInfo
}

func NewTraceGenerator(routineID int, exporter *otlptrace.Exporter, resGen *ResourceGenerator, cfg *config.Config) *TraceGenerator {
	resource, serviceInfo := resGen.GenerateResource()
	if resource == nil {
		log.Printf("Goroutine %d: Resource 생성 실패", routineID)
		return nil
	}

	spanGen := NewSpanGenerator(serviceInfo.ServiceType, cfg)

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter), sdktrace.WithResource(resource))

	return &TraceGenerator{
		routineID:   routineID,
		exporter:    exporter,
		resGen:      resGen,
		tp:          tp,
		spanGen:     spanGen,
		serviceInfo: serviceInfo,
	}
}

func (tg *TraceGenerator) Start(mainCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer tg.Shutdown()

	tracer := tg.tp.Tracer(fmt.Sprintf("otel-generator-periodic-worker-%d", tg.routineID))
	tg.spanGen.tracer = tracer

	ticker := time.NewTicker(time.Duration(traceIntervalSeconds) * time.Second)
	defer ticker.Stop()
	log.Printf("Goroutine %d: Resource(%s) Trace 전송 시작 (간격: %d초)", tg.routineID, tg.serviceInfo.String(), traceIntervalSeconds)

	for {
		select {
		case <-mainCtx.Done():
			log.Printf("Goroutine %d: 종료 신호 수신 (Resource: %s). Trace 전송 중단.", tg.routineID, tg.serviceInfo.String())
			return

		case <-ticker.C:
			tg.spanGen.GenerateTrace(mainCtx)
			//spanName := "periodic-simulated-task"
			//taskCtx, rootSpan := tracer.Start(mainCtx, spanName)
			//tg.spanGen.setPopulateSpanAttributes(rootSpan)
			//time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)
			//
			//_, childSpan := tracer.Start(taskCtx, "sub-task-in-periodic-job")
			//time.Sleep(time.Duration(20+rand.Intn(50)) * time.Millisecond)
			//childSpan.End()

			//log.Printf("Goroutine %d: Resource(%s) - Trace (%s) 전송 완료.", tg.routineID, tg.serviceInfo.String(), rootSpan.SpanContext().TraceID().String())
		}
	}
}

func (tg *TraceGenerator) Shutdown() {
	tpShutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := tg.tp.Shutdown(tpShutdownCtx); err != nil {
		log.Printf("Goroutine %d: TracerProvider Shutdown 실패: %v", tg.routineID, err)
	} else {
		log.Printf("Goroutine %d: TracerProvider Shutdown 완료", tg.routineID)
	}
}
