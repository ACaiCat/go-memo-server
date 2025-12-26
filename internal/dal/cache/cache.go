package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var (
	ErrCacheMiss     = errors.New("cache: key not found")
	ErrSerialization = errors.New("cache: serialization error")
	ErrRedis         = errors.New("cache: redis operation failed")
)

type Cache struct {
	category string
	redis    *redis.Client
}

func NewCache(redis *redis.Client, category string) *Cache {
	return &Cache{
		category: category,
		redis:    redis,
	}
}

func (p *Cache) Get(key string, value any) error {
	cacheKey := fmt.Sprintf("%s:%s", p.category, key)
	jsonCache, err := p.redis.Get(cacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrCacheMiss
		}
		hlog.Errorf("%s: %v", ErrSerialization, err)
		return ErrRedis
	}

	if jsonCache == "null" {
		value = nil
		return nil
	}

	if err := json.Unmarshal([]byte(jsonCache), value); err != nil {
		hlog.Errorf("%s: %v", ErrSerialization, err)
		return ErrSerialization
	}

	return nil
}

func (p *Cache) SetEmpty(key string, ttl time.Duration) error {
	cacheKey := fmt.Sprintf("%s:%s", p.category, key)
	_, err := p.redis.Set(cacheKey, "null", ttl).Result()

	if err != nil {
		hlog.Errorf("%s: %v", ErrSerialization, err)
		return ErrRedis
	}

	return nil
}

func (p *Cache) Set(key string, value any, ttl time.Duration) error {
	cacheKey := fmt.Sprintf("%s:%s", p.category, key)

	jsonCache, err := json.Marshal(value)
	if err != nil {
		hlog.Errorf("%s: %v", ErrSerialization, err)
		return ErrSerialization
	}

	_, err = p.redis.Set(cacheKey, string(jsonCache), ttl).Result()
	if err != nil {
		hlog.Errorf("%s: %v", ErrSerialization, err)
		return ErrRedis
	}
	return nil
}
func (p *Cache) Delete(keys ...string) error {

	for _, key := range keys {
		cacheKey := fmt.Sprintf("%s:%s", p.category, key)
		_, err := p.redis.Del(cacheKey).Result()
		if err != nil {
			hlog.Errorf("%s: %v", ErrSerialization, err)
			return ErrRedis
		}
	}
	return nil
}
