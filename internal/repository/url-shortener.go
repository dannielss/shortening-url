package repository

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func (r *Repo) StoreShortURL(shortURL, longURL string) error {
	err := r.cache.Set(context.Background(), shortURL, longURL, 0).Err()
	if err != nil {
		return fmt.Errorf("could not store short URL: %w", err)
	}
	return nil
}

func (r *Repo) GetLongURL(shortURL string) (string, error) {
	longURL, err := r.cache.Get(context.Background(), shortURL).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("short URL not found")
	} else if err != nil {
		return "", fmt.Errorf("failed to get long URL: %v", err)
	}

	return longURL, nil
}

func (r *Repo) GetShortURL(longURL string) string {
	shortUrl, err := r.cache.Get(context.Background(), longURL).Result()

	if err != nil {
		return ""
	}

	return shortUrl
}
