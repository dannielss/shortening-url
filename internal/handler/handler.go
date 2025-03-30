package handler

import "github.com/danniels/shortening-url/internal/usecase"

type Handler struct {
	uc usecase.IUsecase
}

func NewHandler(uc usecase.IUsecase) *Handler {
	return &Handler{uc: uc}
}
