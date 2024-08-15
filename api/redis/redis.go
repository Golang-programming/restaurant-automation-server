package redis

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var (
	redisClient *redis.Client
	ctx         context.Context
)

func init() {
	ctx = context.Background()

	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func Set(key string, value interface{}) error {
	return redisClient.Set(ctx, key, value, 0).Err()
}

func Get(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

func HSet(key string, fields map[string]interface{}) error {
	return redisClient.HSet(ctx, key, fields).Err()
}

func SAdd(key string, members ...interface{}) error {
	return redisClient.SAdd(ctx, key, members...).Err()
}

func SRem(key string, members ...interface{}) error {
	return redisClient.SRem(ctx, key, members...).Err()
}

func SMembers(key string) ([]string, error) {
	return redisClient.SMembers(ctx, key).Result()
}
