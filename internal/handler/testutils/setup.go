package testutils

import (
	"github.com/danniels/shortening-url/internal/handler"
	mocks "github.com/danniels/shortening-url/internal/mocks/repository"
	"github.com/danniels/shortening-url/internal/usecase"
)

func SetupTest() (*handler.Handler, *mocks.MockRepo) {
	repo := new(mocks.MockRepo)
	uc := usecase.NewUsecase(repo)
	h := handler.NewHandler(uc)

	return h, repo
}
