package main

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"humaversion/service"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// setup handlers and routes
	router := mux.NewRouter()
	service.Service(router)

	// setup server
	server := &http.Server{
		Addr:    ":5000",
		Handler: router,
	}

	log.Println("starting http server...")
	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("http server error: %v\n", err)
		}
		log.Println("stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("http server shutdown error: %v\n", err)
	}
	log.Println("graceful http server shutdown.")
}
