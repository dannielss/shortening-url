package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danniels/shortening-url/docs"
	"github.com/danniels/shortening-url/internal/config"
	"github.com/danniels/shortening-url/internal/handler"
	"github.com/danniels/shortening-url/internal/logger"
	"github.com/danniels/shortening-url/internal/middleware"
	"github.com/danniels/shortening-url/internal/repository"
	"github.com/danniels/shortening-url/internal/router"
	"github.com/danniels/shortening-url/internal/usecase"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

func init() {
	prometheus.MustRegister(middleware.HttpRequestsTotal)
	prometheus.MustRegister(middleware.HttpRequestDuration)
}

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	logger.NewLogger()
	defer logger.Logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Logger.Fatal("Error loading .env file")
	}

	redis_addr := os.Getenv("REDIS_ADDR")

	redis := config.NewCacheClient(redis_addr, "", 0)
	repo := repository.NewRepo(redis)
	uc := usecase.NewUsecase(repo)
	h := handler.NewHandler(uc)
	r := router.SetupRoutes(h)

	srv := &http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	go func() {
		logger.Logger.Info("Server is running on port 3000")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Logger.Fatal("Server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	logger.Logger.Info("Shutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Logger.Info("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		logger.Logger.Fatal("Server forced to shutdown", zap.Error(err))
	}

	logger.Logger.Info("Server exited gracefully")
}
