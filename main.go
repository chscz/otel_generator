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
)

func main() {
	mainCtx, cancelSignal := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancelSignal()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config:%v", err)
	}

	e := exporter.NewExporter(mainCtx, cfg.CollectorURL)
	defer e.Shutdown()

	resourceGenerator := generator.NewResource(cfg.Services)

	var wg sync.WaitGroup
	for i := 0; i < cfg.GoroutineCount; i++ {
		wg.Add(1)
		tg := generator.NewTraceGenerator(i, e.Exp, resourceGenerator, cfg)
		go tg.Start(mainCtx, &wg)
	}

	<-mainCtx.Done()
	log.Println("애플리케이션 종료 신호 수신. 모든 Goroutine 이 종료될 때까지 대기 중...")

	wg.Wait()
	log.Println("모든 Goroutine 정상 종료됨.")
	log.Println("애플리케이션 종료.")
}
