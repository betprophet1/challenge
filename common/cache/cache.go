package cache

import (
	"context"
	"fmt"

	"project/common/cache/client"
)

type GetClientFunc func() client.Client
type GetRateLimiterFunc func() client.RateLimiter
type SimpleLockFunc func(ctx context.Context, key string, params ...any) (client.Locker, error)

var GetClient GetClientFunc = func() client.Client {
	return client.GetRedisClient()
}

var GetRateLimiter GetRateLimiterFunc = func() client.RateLimiter {
	return client.GetRedisLimiter()
}

var LockSimple SimpleLockFunc = func(ctx context.Context, key string, params ...any) (client.Locker, error) {
	if len(params) > 0 {
		key = fmt.Sprintf(key, params...)
	}
	locker := client.GetRedisLocker()
	return locker.Lock(ctx, key)
}
