package usecase

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/mohammaderm/todoMicroService/todoService/internal/models"
)

type (
	todoCache struct {
		redis *redis.Client
	}

	todoCacheInterface interface {
		// todo
		setAll(ctx context.Context, key string, value *[]models.Todo, ttl time.Duration) error
		getAll(ctx context.Context, key string) (*[]models.Todo, error)
		deleteAll(ctx context.Context) error

		// categrory
		setAllCat(ctx context.Context, key string, value *[]models.Category, ttl time.Duration) error
		getAllCat(ctx context.Context, key string) (*[]models.Category, error)
	}
)

func NewCache(redis *redis.Client) todoCacheInterface {
	return &todoCache{
		redis: redis,
	}
}

// categrory
func (c todoCache) getAllCat(ctx context.Context, key string) (*[]models.Category, error) {
	value, err := c.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, err
	}
	var categorys *[]models.Category
	err = json.Unmarshal([]byte(value), &categorys)
	if err != nil {
		return nil, err
	}
	return categorys, nil
}

func (c todoCache) setAllCat(ctx context.Context, key string, value *[]models.Category, ttl time.Duration) error {
	marshalValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.redis.Set(ctx, key, marshalValue, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

// todo
func (c todoCache) setAll(ctx context.Context, key string, value *[]models.Todo, ttl time.Duration) error {
	marshalValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = c.redis.Set(ctx, key, marshalValue, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (c todoCache) getAll(ctx context.Context, key string) (*[]models.Todo, error) {
	value, err := c.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var job *[]models.Todo
	err = json.Unmarshal([]byte(value), &job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

func (c todoCache) deleteAll(ctx context.Context) error {
	err := c.redis.FlushDB(ctx).Err()
	if err != nil {
		return err
	}
	return nil
}
