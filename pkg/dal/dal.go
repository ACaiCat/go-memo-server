package dal

import (
	"github.com/ACaiCat/memo/pkg/dal/cache"
	"github.com/ACaiCat/memo/pkg/dal/db"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Cache *redis.Client
)

func InitDal() {
	DB = db.InitPostgre()
	Cache = cache.InitRedis()
}
