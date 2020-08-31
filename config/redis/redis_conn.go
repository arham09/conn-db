package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func Connect(addr string, password string) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	_, err := conn.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}

	return conn, nil
}
