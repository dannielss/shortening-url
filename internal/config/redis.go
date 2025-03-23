package config

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewCacheClient(redisAddr, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis!")
	return client
}
