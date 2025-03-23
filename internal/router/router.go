package router

import (
	"github.com/danniels/shortening-url/internal/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/shorten", h.ShortUrl)
	r.GET("/:shorted_url", h.UrlRedirect)

	return r
}
