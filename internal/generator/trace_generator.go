package generator

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const traceIntervalSeconds = 5

type TraceGenerator struct {
	routineID int
	exporter  *otlptrace.Exporter
	resGen    *ResourceGenerator
	tp        *sdktrace.TracerProvider
	//Span     *SpanGenerator
	//Platform attrresource.PlatformType
}

func NewTraceGenerator(routineID int, exporter *otlptrace.Exporter, resGen *ResourceGenerator) *TraceGenerator {
	return &TraceGenerator{
		routineID: routineID,
		exporter:  exporter,
		resGen:    resGen,
	}
}

func (tg *TraceGenerator) Start(parentCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	resource, serviceInfo := tg.resGen.GenerateResource()
	if resource == nil {
		log.Printf("Goroutine %d: Resource 생성 실패", tg.routineID)
		return
	}
	log.Printf("Goroutine %d: Resource(%s)로 TracerProvider 설정 중...", tg.routineID, serviceInfo.String())

	tg.tp = sdktrace.NewTracerProvider(sdktrace.WithBatcher(tg.exporter), sdktrace.WithResource(resource))
	defer tg.Shutdown()

	tracer := tg.tp.Tracer(fmt.Sprintf("otel-generator-periodic-worker-%d", tg.routineID))
	spanGenerator := NewSpanGeneratorWithTracer(tracer, serviceInfo.Platform)

	ticker := time.NewTicker(time.Duration(traceIntervalSeconds) * time.Second)
	defer ticker.Stop()
	log.Printf("Goroutine %d: Resource(%s) Trace 전송 시작 (간격: %d초)", tg.routineID, serviceInfo.String(), traceIntervalSeconds)

	for {
		select {
		case <-parentCtx.Done():
			log.Printf("Goroutine %d: 종료 신호 수신 (Resource: %s). Trace 전송 중단.", tg.routineID, serviceInfo.String())
			return
		case <-ticker.C:
			spanName := "periodic-simulated-task"
			taskCtx, rootSpan := tracer.Start(parentCtx, spanName)
			spanGenerator.PopulateSpanAttributes(rootSpan)
			time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)

			_, childSpan := tracer.Start(taskCtx, "sub-task-in-periodic-job")
			time.Sleep(time.Duration(20+rand.Intn(50)) * time.Millisecond)
			childSpan.End()

			rootSpan.End()
			log.Printf("Goroutine %d: Resource(%s) - Trace (%s) 전송 완료.", tg.routineID, serviceInfo.String(), rootSpan.SpanContext().TraceID().String())
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
