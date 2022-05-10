package main

import (
	"L0/internal/repository"
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	connStr := os.Getenv("PG_CONNSTR")

	repo, err := repository.New(connStr)
	service := service.New(&repo)
	router := handlers.New(&service)

	//context and cancel function for server
	srvCtx, srvCancelFunc := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		//blocking until signal caught
		sig := <-c
		log.Printf("%v signal caught, force quit in 30 second", sig)
		shutdownCtx, _ := context.WithTimeout(srvCtx, time.Second*30)

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out, forcing exit")
			}
		}()
		err := router.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal("error shutting down the server:", err)
		}
		srvCancelFunc()
	}()

	err := http.ListenAndServe("8080", router)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server stopped unexpectidly with error: %v", err)
	}

	<-srvCtx.Done()
}
