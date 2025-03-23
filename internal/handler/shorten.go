package handler

import (
	"net/http"

	"github.com/danniels/shortening-url/internal/domain"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ShortUrl(ctx *gin.Context) {
	request := &domain.ShortUrlRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	shortenedURL, err := h.uc.ShortURL(request.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not shorten URL"})
		return
	}

	ctx.JSON(http.StatusOK, shortenedURL)
}
