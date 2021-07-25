package cache

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Client *redis.Client
}

func GetRedisCache(port int) *RedisCache {
	addr := "localhost:" + strconv.Itoa(port)
	redisClient := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalln(err)
	}
	return &RedisCache{Client: redisClient}
}

func (c *RedisCache) Set(key string, value string) error {
	err := c.Client.Set(context.Background(), key, value, 30*time.Minute).Err()
	return err
}

func (c *RedisCache) Get(key string) (string, error) {
	value, err := c.Client.Get(context.Background(), key).Result()
	return value, err
}
