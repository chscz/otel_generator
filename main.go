package main

import (
	"context"
	"fmt"
	"log"
	"math/rand" // 추가: 작업 시뮬레이션 시간 랜덤화
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"otel-generator/internal/exporter"
	generator "otel-generator/internal/generator"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	numGoroutines        = 3
	traceIntervalSeconds = 5
)

func main() {
	mainCtx, cancelSignal := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancelSignal()

	exp := exporter.NewExporter(mainCtx, "http://localhost:4318/v1/traces")

	defer func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := exp.Shutdown(shutdownCtx); err != nil {
			log.Printf("failed to shutdown exporter: %v", err)
		} else {
			log.Println("success shutdown exporter!!!")
		}
	}()

	resourceGenerator := generator.NewResource()

	var wg sync.WaitGroup
	log.Printf("%d개의 리소스(Goroutine)가 각각 Trace를 주기적으로 전송 시작...", numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(routineID int, parentCtx context.Context) {
			defer wg.Done()

			resource := resourceGenerator.GenerateResource()
			if resource == nil {
				log.Printf("Goroutine %d: Resource 생성 실패", routineID)
				return
			}
			log.Printf("Goroutine %d: Resource(%s)로 TracerProvider 설정 중...", routineID, resource.Attributes())

			tp := sdktrace.NewTracerProvider(
				sdktrace.WithBatcher(exp),
				sdktrace.WithResource(resource),
			)

			defer func() {
				tpShutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				if err := tp.Shutdown(tpShutdownCtx); err != nil {
					log.Printf("Goroutine %d: TracerProvider Shutdown 실패 (Resource: %s): %v", routineID, resource.Attributes(), err)
				} else {
					log.Printf("Goroutine %d: TracerProvider Shutdown 완료 (Resource: %s)", routineID, resource.Attributes())
				}
			}()

			tracer := tp.Tracer(fmt.Sprintf("otel-generator-periodic-worker-%d", routineID))
			spanGenerator := generator.NewSpanGeneratorWithTracer(tracer)

			ticker := time.NewTicker(time.Duration(traceIntervalSeconds) * time.Second)
			defer ticker.Stop()

			log.Printf("Goroutine %d: Resource(%s) Trace 전송 시작 (간격: %d초)", routineID, resource.Attributes(), traceIntervalSeconds)

			// 4.e. 주기적 Trace 생성 루프 (애플리케이션 종료 신호 전까지)
			for {
				select {
				case <-parentCtx.Done(): // mainCtx가 취소되었는지 (종료 신호) 확인
					log.Printf("Goroutine %d: 종료 신호 수신 (Resource: %s). Trace 전송 중단.", routineID, resource.Attributes())
					return // Goroutine 종료

				case <-ticker.C: // Ticker가 울릴 때마다 Trace 생성
					spanName := "periodic-simulated-task"
					// 각 Trace 작업에 대한 Context 생성 (parentCtx를 부모로 하여 취소 전파 가능)
					taskCtx, rootSpan := tracer.Start(parentCtx, spanName)

					spanGenerator.PopulateSpanAttributes(rootSpan) // Span에 속성 추가

					log.Printf("Goroutine %d: Resource(%s) - Trace (%s) 생성 중...", routineID, resource.Attributes(), rootSpan.SpanContext().TraceID().String())
					// 간단한 작업 시뮬레이션 (랜덤 시간)
					time.Sleep(time.Duration(50+rand.Intn(100)) * time.Millisecond)

					// 자식 Span 예시 (선택 사항)
					_, childSpan := tracer.Start(taskCtx, "sub-task-in-periodic-job")
					time.Sleep(time.Duration(20+rand.Intn(50)) * time.Millisecond)
					childSpan.End()

					rootSpan.End() // 루트 Span 종료
					log.Printf("Goroutine %d: Resource(%s) - Trace (%s) 전송 완료.", routineID, resource.Attributes(), rootSpan.SpanContext().TraceID().String())
				}
			}
		}(i, mainCtx)
	}

	// 5. mainCtx가 취소될 때까지 (즉, 종료 신호 수신) 대기
	<-mainCtx.Done()
	log.Println("애플리케이션 종료 신호 수신 (예: Ctrl+C). 모든 Goroutine이 종료될 때까지 대기 중...")

	// 6. 모든 Goroutine이 정상적으로 종료될 때까지 대기
	wg.Wait()
	log.Println("모든 Goroutine 정상 종료됨.")
	log.Println("애플리케이션 종료.")
}
