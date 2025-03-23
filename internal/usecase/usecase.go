package usecase

import (
	"github.com/danniels/shortening-url/internal/domain"
)

type usecase struct {
	// repo *repository.Repo
}

type Usecase interface {
	ShortURL(url string) (*domain.UrlMapping, error)
	GetOriginalURL(shortedUrl string) (string, error)
}

// func NewUsecase(repo *repository.Repo) *usecase {
// 	return &usecase{repo: repo}
// }

func NewUsecase() *usecase {
	return &usecase{}
}
