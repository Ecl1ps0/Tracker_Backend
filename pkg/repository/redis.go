package repository

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisDB(dsn string) (*redis.Client, error) {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{client: client}
}

func (r *RedisRepository) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	serializedValue, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, serializedValue, expiration).Err()
}

func (r *RedisRepository) Get(ctx context.Context, key string, dest interface{}) error {
	serializedValue, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(serializedValue), dest)
}

func (r *RedisRepository) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
