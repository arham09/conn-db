package redis

import (
	"github.com/go-redis/redis/v8"
)

func Connect(addr string, password string) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	return conn, nil
}
