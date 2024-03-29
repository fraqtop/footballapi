package main

import (
	"context"
	"github.com/fraqtop/footballapi/internal/config"
	"github.com/fraqtop/footballapi/internal/container"
	"github.com/fraqtop/footballapi/internal/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	interruptContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	container.Init()

	go func() {
		if err := server.Serve(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-interruptContext.Done()

	log.Println("graceful shutting down")

	serverContext, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := server.Destroy(serverContext); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
