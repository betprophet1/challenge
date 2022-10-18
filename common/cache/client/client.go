package client

import (
	"context"
	"time"

	"github.com/go-redis/redis_rate/v9"
)

type (
	Options struct {
		Timeout time.Duration
	}

	Client interface {
		Exists(ctx context.Context, key string) (bool, error)
		Get(ctx context.Context, key string, result interface{}) error
		Set(ctx context.Context, key string, value interface{}, timeout time.Duration) error
		SetEx(ctx context.Context, key string, value interface{}, options Options) error
		Delete(ctx context.Context, key string) error
		Incr(ctx context.Context, key string) error
		Decr(ctx context.Context, key string) error
	}

	RateLimiter interface {
		AllowPerSecond(ctx context.Context, key string, rate int) (*redis_rate.Result, error)
	}

	Locker interface {
		Lock(ctx context.Context, key string) (Locker, error)
		Unlock(ctx context.Context) (bool, error)
	}
)
