package router

import (
	"github.com/danniels/shortening-url/internal/handler"
	"github.com/danniels/shortening-url/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.MetricsMiddleware())

	r.POST("/shorten", h.ShortUrl)
	r.GET("/:shorted_url", h.UrlRedirect)

	// Expose Prometheus metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
