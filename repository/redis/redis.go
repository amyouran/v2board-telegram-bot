package redis

import (
	"context"
	"sync"
	"time"

	"v2board-telegram-bot/configs"

	redis "github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
)

const Nil = redis.Nil

type Pipeliner interface {
	redis.Pipeliner
}

func GetRedisClient() *redis.Client {
	once.Do(func() {
		cfg := configs.Get().Redis
		client = redis.NewClient(&redis.Options{
			Addr:         cfg.Addr,
			Password:     cfg.Password,
			DB:           cfg.DB,
			MaxRetries:   3,
			DialTimeout:  time.Second * 5,
			ReadTimeout:  time.Second * 20,
			WriteTimeout: time.Second * 20,
			PoolSize:     50,
			MinIdleConns: 2,
			PoolTimeout:  time.Minute,
		})

		if err := client.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}

	})

	return client
}
