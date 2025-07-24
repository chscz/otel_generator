package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"otel-generator/internal/config"
	"otel-generator/internal/exporter"
	"otel-generator/internal/generator"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	mainCtx, cancelSignal := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancelSignal()

	cfg, err := config.LoadConfig("./config.yaml")
	if err != nil {
		log.Fatalf("failed to load config:%v", err)
	}

	e := exporter.NewExporter(mainCtx, cfg.CollectorURL)
	defer e.Shutdown()
	batchProcessor := sdktrace.NewBatchSpanProcessor(e.Exp)
	defer batchProcessor.Shutdown(context.Background())

	resourceGenerator := generator.NewResourceGenerator(cfg.Services, cfg.ResourceAttributes)

	var wg sync.WaitGroup
	for i := 0; i < cfg.GoRoutineCount; i++ {
		wg.Add(1)
		tg, err := generator.NewTraceGenerator(mainCtx, i, batchProcessor, resourceGenerator, cfg)
		if err != nil {
			log.Printf("failed to create trace generator:%v", err)
			wg.Done()
			continue
		}
		if cfg.GenerateOption.UseDynamicInterval {
			go tg.StartDynamicInterval(mainCtx, &wg)
		} else {
			go tg.StartFixedInterval(mainCtx, &wg)
		}
	}

	<-mainCtx.Done()
	log.Println("애플리케이션 종료 신호 수신. 모든 Goroutine 이 종료될 때까지 대기 중...")

	wg.Wait()
	log.Println("success graceful shutdown!! 모든 Goroutine 정상 종료됨.")
	log.Println("애플리케이션 종료.")
}
