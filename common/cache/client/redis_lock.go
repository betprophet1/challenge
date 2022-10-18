package client

import (
	"context"

	"github.com/go-redsync/redsync/v4"
)

type RedisMutex struct {
	mutex *redsync.Mutex
}

func (m *RedisMutex) Lock(ctx context.Context, key string) (Locker, error) {
	m.mutex = redSyncPool.NewMutex(key)
	if err := m.mutex.LockContext(ctx); err != nil {
		return nil, err
	}
	return m, nil
}

func (m RedisMutex) Unlock(ctx context.Context) (bool, error) {
	return m.mutex.UnlockContext(ctx)
}

func GetRedisLocker() Locker {
	return &RedisMutex{}
}
