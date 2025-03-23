package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danniels/shortening-url/internal/config"
	"github.com/danniels/shortening-url/internal/handler"
	"github.com/danniels/shortening-url/internal/repository"
	"github.com/danniels/shortening-url/internal/router"
	"github.com/danniels/shortening-url/internal/usecase"
)

func main() {
	redis := config.NewCacheClient("localhost:6379", "", 0)

	repo := repository.NewRepo(redis)

	uc := usecase.NewUsecase(repo)

	h := handler.NewHandler(uc)

	r := router.SetupRoutes(h)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		log.Println("Server is running on port 3000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
