package handler

import (
	"net/http"

	"github.com/danniels/shortening-url/internal/domain"
	"github.com/gin-gonic/gin"
)

// ShortUrl godoc
//
//	@Summary		Shorten a URL
//	@Description	Takes a long URL and returns a shortened version
//	@Tags			URL
//	@Accept			json
//	@Produce		json
//	@Param			request	body	domain.ShortUrlRequest	true	"URL to shorten"
//	@Success		200	{object} domain.UrlMapping
//	@Router			/shorten [post]
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
