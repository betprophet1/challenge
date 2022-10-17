package client

import (
	"context"
	"time"
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
)
