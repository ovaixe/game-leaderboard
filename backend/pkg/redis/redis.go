package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisClient represents the Redis client
type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient initializes and returns a new Redis client
func NewRedisClient() (*RedisClient, error) {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if redisHost == "" || redisPort == "" {
		return nil, fmt.Errorf("REDIS_HOST and REDIS_PORT environment variables must be set")
	}

	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, fmt.Errorf("could not connect to Redis: %w", err)
	}

	return &RedisClient{Client: rdb}, nil
}

// Close closes the Redis client connection
func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}
