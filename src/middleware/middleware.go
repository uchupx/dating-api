package middleware

import "github.com/uchupx/dating-api/pkg/database/redis"

type Middleware struct {
	Redis *redis.Redis
}
