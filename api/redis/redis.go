package redis

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Update this with your Redis server address
		DB:   0,                // Default DB
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}

// SaveToRedis saves data to Redis with the specified key
func SaveToRedis(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return redisClient.Set(ctx, key, data, 0).Err()
}

// GetFromRedis retrieves data from Redis by key
func GetFromRedis(key string, dest interface{}) error {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil
		}
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

// DeleteFromRedis removes a key from Redis
func DeleteFromRedis(key string) error {
	return redisClient.Del(ctx, key).Err()
}
