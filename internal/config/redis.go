package config

import (
	"context"

	"github.com/danniels/shortening-url/internal/logger"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func NewCacheClient(redisAddr, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		logger.Logger.Fatal("Could not connect to Redis", zap.Error(err))
	}

	logger.Logger.Info("Successfully connected to Redis!")
	return client
}
