package wager

import (
	"context"
	"fmt"
	"time"

	"project/common/cache/client"
)

var counterkey = "simplebet:buy:[wager<%d>]:count"

// Cache sold counts
type SoldCounter struct {
	WagerID uint64
	Mem     client.Client
}

func (s SoldCounter) Incr(ctx context.Context) error {
	simplekey := fmt.Sprintf(counterkey, s.WagerID)
	ok, err := s.Mem.Exists(ctx, simplekey)
	if err != nil {
		return err
	}
	if ok {
		if err = s.Mem.Incr(ctx, simplekey); err != nil {
			return err
		}
		return nil
	}
	// assume an wager is just valid for 24hour for purchasing
	// Note for futher data handling in memory
	return s.Mem.Set(ctx, simplekey, 1, 24*time.Hour)
}

func (s SoldCounter) GetCount(ctx context.Context) (count uint, err error) {
	simplekey := fmt.Sprintf(counterkey, s.WagerID)
	err = s.Mem.Get(ctx, simplekey, &count)
	return
}

func (s SoldCounter) Decr(ctx context.Context) error {
	simplekey := fmt.Sprintf(counterkey, s.WagerID)
	return s.Mem.Decr(ctx, simplekey)
}
