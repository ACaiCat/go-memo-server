package cache

import (
	"github.com/ACaiCat/memo/pkg/config"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	rdConfig := config.GetConfig().RedisConfig
	cache := redis.NewClient(&redis.Options{
		Addr:     rdConfig.Address,
		Password: rdConfig.Password,
		DB:       rdConfig.DB,
	})

	_, err := cache.Ping().Result()
	if err != nil {
		hlog.Fatalf("Failed to connect redis: %v", err)
	}
	return cache
}
