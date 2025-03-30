package handler_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danniels/shortening-url/internal/handler/testutils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestURLRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	shortedURL := "EAaArVRs"
	longURL := "https://example.com"

	t.Run("Success", func(t *testing.T) {
		h, repo := testutils.SetupTest()
		repo.On("GetLongURL", shortedURL).Return(longURL, nil)

		router := gin.Default()
		router.GET("/:shorted_url", h.UrlRedirect)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", shortedURL), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		repo.AssertExpectations(t)

		assert.Equal(t, http.StatusMovedPermanently, w.Code)
	})

	t.Run("Usecase Error", func(t *testing.T) {
		h, repo := testutils.SetupTest()

		repo.On("GetLongURL", shortedURL).Return("", errors.New("error"))

		router := gin.Default()
		router.GET("/:shorted_url", h.UrlRedirect)

		req, _ := http.NewRequest("GET", fmt.Sprintf("/%s", shortedURL), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		repo.AssertExpectations(t)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "URL not found")
	})
}
