package client

import (
	"context"

	"github.com/go-redis/redis_rate/v9"
)

func (c RedisClient) AllowPerSecond(ctx context.Context, key string, rate int) (*redis_rate.Result, error) {
	return c.limiter.Allow(ctx, key, redis_rate.PerSecond(rate))
}

func GetRedisLimiter() RateLimiter {
	return redisLimiter
}
