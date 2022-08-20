package main

import (
	"context"
	"github.com/fraqtop/footballapi/config"
	"github.com/fraqtop/footballapi/connection"
	"github.com/fraqtop/footballapi/server"
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

	if err := connection.Init(); err != nil {
		log.Fatal(err)
	}
	defer connection.Destroy()

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
