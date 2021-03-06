package caching

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type Caching interface {
	SetItem(ctx context.Context, key string, item interface{}, duration time.Duration) error
	GetItem(ctx context.Context, key string) (string, bool)
	DelItem(ctx context.Context, key string) error
	DelWildItem(ctx context.Context, prefix string) error
}

type redisCaching struct {
	Conn *redis.Client
}

func NewRedisCaching(Conn *redis.Client) Caching {
	return &redisCaching{Conn}
}

func (r *redisCaching) SetItem(ctx context.Context, key string, item interface{}, duration time.Duration) error {
	data, err := json.Marshal(item)

	if err != nil {
		return err
	}

	err = r.Conn.Set(ctx, key, string(data), duration).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisCaching) GetItem(ctx context.Context, key string) (string, bool) {
	value, err := r.Conn.Get(ctx, key).Result()

	if err != nil {
		return "", false
	}

	return value, true
}

func (r *redisCaching) DelItem(ctx context.Context, key string) error {
	err := r.Conn.Del(ctx, key).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *redisCaching) DelWildItem(ctx context.Context, prefix string) error {
	iter := r.Conn.Scan(ctx, 0, prefix, 0).Iterator()

	for iter.Next(ctx) {
		err := r.Conn.Del(ctx, iter.Val()).Err()

		if err != nil {
			return err
		}
	}

	if err := iter.Err(); err != nil {
		panic(err)
	}

	return nil
}
