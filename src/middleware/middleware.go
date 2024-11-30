package middleware

import "github.com/uchupx/dating-api/pkg/database/redis"

type Middleware struct {
	Redis *redis.Redis
}

type Config struct {
	Redis *redis.Redis
}

func New(conf Config) *Middleware {
	return &Middleware{
		Redis: conf.Redis,
	}
}
