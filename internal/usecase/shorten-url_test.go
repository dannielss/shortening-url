package usecase

import (
	"errors"
	"testing"

	mocks "github.com/danniels/shortening-url/internal/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func TestShortURL(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	uc := NewUsecase(mockRepo)

	originalURL := "https://example.com"
	shortURL := "EAaArVRs"

	mockRepo.On("GetLongURL", shortURL).Return("", nil)
	mockRepo.On("StoreShortURL", shortURL, originalURL).Return(nil)

	result, err := uc.ShortenURL(originalURL)

	assert.NoError(t, err)
	assert.Equal(t, originalURL, result.LongURL)
	assert.Equal(t, shortURL, result.ShortURL)

	mockRepo.AssertExpectations(t)
}

func TestShortURLInRedis(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	uc := NewUsecase(mockRepo)

	originalURL := "https://example.com"
	shortURL := "EAaArVRs"

	mockRepo.On("GetLongURL", shortURL).Return("https://example.com", nil)

	result, err := uc.ShortenURL(originalURL)

	assert.NoError(t, err)
	assert.Equal(t, originalURL, result.LongURL)
	assert.Equal(t, shortURL, result.ShortURL)

	mockRepo.AssertExpectations(t)
}

func TestShortURLError(t *testing.T) {
	mockRepo := new(mocks.MockRepo)
	uc := NewUsecase(mockRepo)

	originalURL := "https://example.com"
	shortURL := "EAaArVRs"

	mockRepo.On("GetLongURL", shortURL).Return("", nil)
	mockRepo.On("StoreShortURL", shortURL, originalURL).Return(errors.New("error"))

	_, err := uc.ShortenURL(originalURL)

	assert.Error(t, err)
}
