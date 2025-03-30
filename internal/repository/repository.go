package repository

import (
	"github.com/go-redis/redis/v8"
)

type Repo struct {
	cache *redis.Client
}

type IRepository interface {
	StoreShortURL(shortURL, longURL string) error
	GetLongURL(shortURL string) (string, error)
	GetShortURL(longURL string) string
}

func NewRepo(cache *redis.Client) *Repo {
	return &Repo{cache: cache}
}
