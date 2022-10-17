package client

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"time"

	"project/common/singleton"
	"project/config"

	"github.com/go-redis/redis/v8"
)

var (
	ErrNotFound = errors.New("cache: data not found")
	ErrSetValue = errors.New("cache: cannot set to value")
)

var redisClient Client

func InitRedis() {
	singleRedis := singleton.New(func() interface{} {
		client := redis.NewClient(&redis.Options{
			Addr: os.Getenv(config.SERVICE_REDIS_DNS),
			DB:   0,
		})
		return &RedisClient{client}
	})
	redisClient = singleRedis.Get().(*RedisClient)
}

type RedisClient struct {
	client *redis.Client
}

func (rc *RedisClient) Get(ctx context.Context, key string, result interface{}) error {
	cmd := rc.client.Get(ctx, key)
	if cmd.Err() == redis.Nil {
		return ErrNotFound
	}
	bytes, err := cmd.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, &result)
}

func (rc *RedisClient) Exists(ctx context.Context, key string) (bool, error) {
	cmd := rc.client.Exists(ctx, key)
	result, err := cmd.Result()
	return result > 0, err
}

func (rc *RedisClient) Set(
	ctx context.Context,
	key string, value interface{}, timeout time.Duration,
) error {
	return rc.SetEx(ctx, key, value, Options{Timeout: timeout})
}

func (rc *RedisClient) SetEx(
	ctx context.Context,
	key string, value interface{}, options Options,
) error {
	valueBytes, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return rc.client.Set(ctx, key, valueBytes, options.Timeout).Err()
}

func (rc *RedisClient) Delete(ctx context.Context, key string) error {
	return rc.client.Del(ctx, key).Err()
}

func (rc *RedisClient) Incr(ctx context.Context, key string) error {
	return rc.client.Incr(ctx, key).Err()
}

func (rc *RedisClient) Decr(ctx context.Context, key string) error {
	return rc.client.Decr(ctx, key).Err()
}

func GetRedisClient() Client {
	return redisClient
}
