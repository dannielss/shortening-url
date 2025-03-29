package usecase

import (
	"github.com/danniels/shortening-url/internal/domain"
	"github.com/danniels/shortening-url/internal/repository"
)

type usecase struct {
	repo repository.Repository
}

type Usecase interface {
	ShortURL(url string) (*domain.UrlMapping, error)
	GetOriginalURL(shortedUrl string) (string, error)
}

func NewUsecase(repo repository.Repository) *usecase {
	return &usecase{repo: repo}
}
