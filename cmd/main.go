package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/PongponZ/go-book-store/config"
	"github.com/PongponZ/go-book-store/internal/bootstrap"
)

func main() {
	cfg := config.Get()

	app, err := bootstrap.Initialize(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := app.HTTPServer.Start(fmt.Sprintf(":%d", cfg.HTTP.Port)); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	<-exit

	log.Println("shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Fatalf("server forced to shutdown app: %v", err)
	}

	log.Println("shutdown gracefully")
}
