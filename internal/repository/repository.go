package repository

import (
	"github.com/go-redis/redis/v8"
)

type Repo struct {
	cache *redis.Client
}

func NewRepo(cache *redis.Client) *Repo {
	return &Repo{cache: cache}
}
