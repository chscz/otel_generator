package generator

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"otel-generator/internal/config"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	traceIntervalMinSeconds = 10
	traceIntervalMaxSeconds = 60
)

type TraceGenerator struct {
	routineID   int
	exporter    *otlptrace.Exporter
	resGen      *ResourceGenerator
	spanGen     *SpanGenerator
	tp          *sdktrace.TracerProvider
	serviceInfo ResourceInfo
}

func NewTraceGenerator(routineID int, exporter *otlptrace.Exporter, resGen *ResourceGenerator, cfg *config.Config) (*TraceGenerator, error) {
	resource, serviceInfo := resGen.GenerateResource()
	if resource == nil {
		return nil, fmt.Errorf("goroutine %d: Resource 생성 실패", routineID)
	}

	spanGen := NewSpanGenerator(serviceInfo.ServiceType, cfg, routineID)

	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exporter), sdktrace.WithResource(resource))

	return &TraceGenerator{
		routineID:   routineID,
		exporter:    exporter,
		resGen:      resGen,
		tp:          tp,
		spanGen:     spanGen,
		serviceInfo: serviceInfo,
	}, nil
}

func (tg *TraceGenerator) StartDynamicInterval(mainCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer tg.Shutdown()

	tracer := tg.tp.Tracer(fmt.Sprintf("otel-generator-periodic-worker-%d", tg.routineID))
	tg.spanGen.tracer = tracer

	localRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	interval := localRand.Intn(traceIntervalMaxSeconds-traceIntervalMinSeconds+1) + traceIntervalMinSeconds
	timer := time.NewTimer(time.Duration(interval) * time.Second)

	defer timer.Stop()
	log.Printf("Goroutine %d: Resource(%s) Trace 전송 시작 (간격: %d초)", tg.routineID, tg.serviceInfo.String(), interval)

	for {
		select {
		case <-mainCtx.Done():
			log.Printf("Goroutine %d: 종료 신호 수신 (Resource: %s). Trace 전송 중단.", tg.routineID, tg.serviceInfo.String())
			return

		case <-timer.C:
			tg.spanGen.GenerateTrace(mainCtx)
			nextInterval := localRand.Intn(traceIntervalMaxSeconds-traceIntervalMinSeconds+1) + traceIntervalMinSeconds
			timer.Reset(time.Duration(nextInterval) * time.Second)
			//log.Printf("Goroutine %d: Resource(%s) - Trace (%s) 전송 완료.", tg.routineID, tg.serviceInfo.String(), rootSpan.SpanContext().TraceID().String())
		}
	}
}

func (tg *TraceGenerator) StartFixedInterval(mainCtx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	defer tg.Shutdown()

	tracer := tg.tp.Tracer(fmt.Sprintf("otel-generator-periodic-worker-%d", tg.routineID))
	tg.spanGen.tracer = tracer

	interval := rand.Intn(traceIntervalMaxSeconds-traceIntervalMinSeconds+1) + traceIntervalMinSeconds
	ticker := time.NewTicker(time.Duration(interval) * time.Second)

	defer ticker.Stop()
	log.Printf("Goroutine %d: Resource(%s) Trace 전송 시작 (간격: %d초)", tg.routineID, tg.serviceInfo.String(), interval)

	for {
		select {
		case <-mainCtx.Done():
			log.Printf("Goroutine %d: 종료 신호 수신 (Resource: %s). Trace 전송 중단.", tg.routineID, tg.serviceInfo.String())
			return

		case <-ticker.C:
			tg.spanGen.GenerateTrace(mainCtx)
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
