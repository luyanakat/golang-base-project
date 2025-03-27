package redis

import (
	"context"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisClient(ctx *context.Context) (*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &RedisClient{
		client: rdb,
		ctx:    *ctx,
	}, nil
}

func (r *RedisClient) SetWithTTL(key string, value interface{}, ttl time.Duration) error {
	err := r.client.Set(r.ctx, key, value, ttl).Err()
	return err
}

func (r *RedisClient) Get(key string) (string, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	return val, err
}

func (r *RedisClient) GetValExist(key string) (string, bool) {
	vs, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return "", false
	}
	return vs, true
}

func (r *RedisClient) Delete(key string) error {
	err := r.client.Del(r.ctx, key).Err()
	return err
}

func (r *RedisClient) DelKeys(keys ...string) error {
	err := r.client.Del(r.ctx, keys...).Err()
	return err
}

func (r *RedisClient) Close() {
	r.client.Close()
}
