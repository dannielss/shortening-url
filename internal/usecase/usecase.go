package usecase

import (
	"github.com/danniels/shortening-url/internal/domain"
	"github.com/danniels/shortening-url/internal/repository"
)

type Usecase struct {
	repo repository.IRepository
}

type IUsecase interface {
	ShortenURL(url string) (*domain.UrlMapping, error)
	GetOriginalURL(shortedUrl string) (string, error)
}

func NewUsecase(repo repository.IRepository) *Usecase {
	return &Usecase{repo: repo}
}
