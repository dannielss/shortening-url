package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// UrlRedirect godoc
//
//	@Summary		Redirect to Long URL
//	@Description	Takes a short URL and redirects to the original long URL
//	@Tags			URL
//	@Accept			json
//	@Produce		json
//	@Param			shorted_url	path	string	true	"Shortened URL"
//	@Success		301	"Redirects to the original URL"
//	@Router			/{shorted_url} [get]
func (h *Handler) UrlRedirect(ctx *gin.Context) {
	shortURL := ctx.Param("shorted_url")

	longURL, err := h.uc.GetOriginalURL(shortURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, longURL)
}
