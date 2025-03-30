package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danniels/shortening-url/internal/domain"
	"github.com/danniels/shortening-url/internal/handler/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShortUrlHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	shortedURL := "EAaArVRs"
	longURL := "https://example.com"

	t.Run("Success", func(t *testing.T) {
		h, repo := testutils.SetupTest()
		repo.On("GetLongURL", shortedURL).Return("", nil)
		repo.On("StoreShortURL", shortedURL, longURL).Return(nil)

		router := gin.Default()
		router.POST("/shorten", h.ShortUrl)

		reqBody, _ := json.Marshal(domain.ShortUrlRequest{URL: longURL})
		req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		repo.AssertExpectations(t)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "short_url")
	})

	t.Run("Bind Error", func(t *testing.T) {
		h, _ := testutils.SetupTest()
		router := gin.Default()
		router.POST("/shorten", h.ShortUrl)

		invalidBody := map[string]string{}
		reqBody, _ := json.Marshal(invalidBody)
		req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), "Invalid request body")
	})

	t.Run("Usecase Error", func(t *testing.T) {
		h, repo := testutils.SetupTest()
		shortedURL := "EAaArVRs"
		longURL := "https://example.com"

		repo.On("GetLongURL", shortedURL).Return("", nil)
		repo.On("StoreShortURL", shortedURL, longURL).Return(errors.New("error"))

		router := gin.Default()
		router.POST("/shorten", h.ShortUrl)

		reqBody, _ := json.Marshal(domain.ShortUrlRequest{URL: longURL})
		req, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		repo.AssertExpectations(t)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}
