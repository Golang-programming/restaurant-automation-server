package redis

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var Client *redis.Client

func init() {
	fmt.Println("——————————————————————————————————————-")

	Client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})
	fmt.Println("€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€€-")

	_, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

func GetRedisClient() *redis.Client {
	return Client
}

func Set(key string, value interface{}, ttl time.Duration) error {
	return Client.Set(ctx, key, value, ttl).Err()
}

func Get(key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

func HSet(key string, fields map[string]interface{}) error {
	return Client.HSet(ctx, key, fields).Err()
}

func SAdd(key string, members ...interface{}) error {
	return Client.SAdd(ctx, key, members...).Err()
}

func SRem(key string, members ...interface{}) error {
	return Client.SRem(ctx, key, members...).Err()
}

func SMembers(key string) ([]string, error) {
	return Client.SMembers(ctx, key).Result()
}

func Keys(pattern string) ([]string, error) {
	return Client.Keys(context.Background(), pattern).Result()
}
