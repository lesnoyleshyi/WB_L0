package main

import (
	"L0/internal/handlers"
	"L0/internal/repository"
	"L0/internal/services"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const restoreGap = 24

func main() {
	repo := repository.New()
	service := services.New(repo)
	handler := handlers.New(service)
	defer func() { _ = (*handler.Nats.Subscription).Unsubscribe() }()
	defer func() { _ = (*handler.Nats.Conn).Close() }()
	router := handler.Routes()
	server := http.Server{Addr: ":8080", Handler: router}
	err := repo.RestoreCache(time.Now().Add(-1 * time.Hour * restoreGap))
	if err != nil {
		log.Fatalf("unable restore cache from database: %s", err)
	}

	//debug function
	go printCache(repo)

	//context and cancel function for server
	srvCtx, srvCancelFunc := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGTERM)

	//non-graceful graceful shutdown
	go func() {
		//blocking until signal caught
		sig := <-c
		log.Printf("%v signal caught, force quit in 30 second", sig)
		shutdownCtx, cf := context.WithTimeout(srvCtx, time.Second*30)
		defer cf()

		go func() {
			<-shutdownCtx.Done()
			if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
				log.Fatal("graceful shutdown timed out, forcing exit")
			}
		}()
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("error shutting down the server: %v", err)
		}
		srvCancelFunc()
	}()

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed unexpectidly with error: %v", err)
	}

	<-srvCtx.Done()
	log.Println("Server stopped safely")
}

func printCache(repo *repository.Repository) {
	for {
		time.Sleep(time.Second * 10)
		fmt.Println(repo.CacheRepo.Storage)
	}
}
