package handler

import "github.com/danniels/shortening-url/internal/usecase"

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}
