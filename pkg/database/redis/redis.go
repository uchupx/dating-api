package redis

import (
	"context"
	"fmt"
	"time"

	goRedis "github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
	MainKey  string `yaml:"main_key"`
}
type Redis struct {
	redis   *goRedis.Client
	mainKey string
}

func (r *Redis) Get(ctx context.Context, key string) (*string, error) {
	result := r.redis.Get(ctx, r.key(key))

	if result.Err() != nil {
		if result.Err() == goRedis.Nil {
			return nil, nil
		}

		return nil, result.Err()
	}

	val := result.Val()

	return &val, nil
}

func (r *Redis) Set(ctx context.Context, key string, val string, ttl *time.Duration) error {
	var duration time.Duration = -1

	if ttl != nil {
		duration = *ttl
	}

	result := r.redis.Set(ctx, r.key(key), val, duration)

	if result.Err() != nil {
		if result.Err() == goRedis.Nil {
			return fmt.Errorf("[Redis] redis nil - failed to set key %s: %w", key, result.Err())
		}

		return result.Err()
	}

	return nil
}

func (r *Redis) key(suffix string) string {
	return r.mainKey + ":" + suffix
}

func Connection(c Config) (*Redis, error) {
	redisClient := goRedis.NewClient(&goRedis.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.Database,
	})

	if err := redisClient.Ping(context.TODO()).Err(); err != nil {
		return nil, fmt.Errorf("failed to ping redis: %w", err)
	}

	return &Redis{
		redis:   redisClient,
		mainKey: c.MainKey,
	}, nil
}
