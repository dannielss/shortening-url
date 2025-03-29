package usecase

import (
	"errors"
	"testing"

	mocks "github.com/danniels/shortening-url/internal/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetOriginalURL(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	uc := NewUsecase(mockRepo)

	originalURL := "https://example.com"
	shortURL := "EAaArVRs"

	mockRepo.On("GetLongURL", shortURL).Return(originalURL, nil)

	longurl, err := uc.GetOriginalURL(shortURL)

	assert.NoError(t, err)
	assert.Equal(t, originalURL, longurl)

	mockRepo.AssertExpectations(t)
}

func TestGetOriginalURLError(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	uc := NewUsecase(mockRepo)

	shortURL := "EAaArVRs"

	mockRepo.On("GetLongURL", shortURL).Return("", errors.New("error"))

	_, err := uc.GetOriginalURL(shortURL)

	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
}
