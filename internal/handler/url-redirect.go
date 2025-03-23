package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RedirectUrl(ctx *gin.Context) {
	shortURL := ctx.Param("shorted_url")

	longURL, err := h.uc.GetOriginalURL(shortURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, longURL)
}
