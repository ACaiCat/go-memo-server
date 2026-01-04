package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/ACaiCat/memo/internal/model"
	"github.com/ACaiCat/memo/pkg/dal/cache"
	"github.com/ACaiCat/memo/pkg/dal/db"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uint) (*model.User, error)
	GetByName(name string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
}

type userRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	cacheClient := cache.NewCache(redis, "user")
	return &userRepository{
		db:    db,
		cache: cacheClient,
	}
}

func (p *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	cacheKey := fmt.Sprintf("id:%d", id)
	if err := p.cache.Get(cacheKey, &user); err == nil {
		hlog.Debugf("Hit user cache: %d\n", id)
		return &user, nil
	}

	if err := p.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := p.cache.SetEmpty(cacheKey, time.Minute*5); err != nil {
				hlog.Errorf("Cache empty user %d failed: %v\n", id, err)
			}
			return nil, nil
		}
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return nil, db.ErrDB
	}

	if err := p.cache.Set(cacheKey, user, time.Minute*30); err != nil {
		hlog.Errorf("Cache user %d failed: %v\n", id, err)
	}

	return &user, nil

}

func (p *userRepository) GetByName(name string) (*model.User, error) {
	var user model.User

	cacheKey := fmt.Sprintf("name:%s", name)
	if err := p.cache.Get(cacheKey, &user); err == nil {
		hlog.Debugf("Hit user cache: %s\n", name)
		return &user, nil
	}

	if err := p.db.Where("name = ?", name).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := p.cache.SetEmpty(cacheKey, time.Minute*5); err != nil {
				hlog.Errorf("Cache empty user %s failed: %v\n", name, err)
			}
			return nil, nil
		}
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return nil, db.ErrDB
	}

	if err := p.cache.Set(cacheKey, user, time.Minute*30); err != nil {
		hlog.Errorf("Cache user %s failed: %v\n", name, err)
	}

	return &user, nil
}

func (p *userRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User

	cacheKey := fmt.Sprintf("email:%s", email)
	if err := p.cache.Get(cacheKey, &user); err == nil {
		hlog.Debugf("Hit user cache: %s\n", email)
		return &user, nil
	}

	if err := p.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := p.cache.SetEmpty(cacheKey, time.Minute*5); err != nil {
				hlog.Errorf("Cache empty user %s failed: %v\n", email, err)
			}
			return nil, nil
		}
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return nil, db.ErrDB
	}

	if err := p.cache.Set(cacheKey, user, time.Minute*30); err != nil {
		hlog.Errorf("Cache user %s failed: %v\n", email, err)
	}

	return &user, nil
}

func (p *userRepository) Create(user *model.User) error {
	if err := p.db.Create(user).Error; err != nil {
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return db.ErrDB
	}
	idCacheKey := fmt.Sprintf("id:%d", user.ID)
	if err := p.cache.Set(idCacheKey, user, time.Minute*30); err != nil {
		hlog.Infof("Cache user %d failed: %v\n", user.ID, err)
	}

	nameCacheKey := fmt.Sprintf("name:%s", user.Name)
	if err := p.cache.Set(nameCacheKey, user, time.Minute*30); err != nil {
		hlog.Infof("Cache user %d failed: %v\n", user.ID, err)
	}

	emailCacheKey := fmt.Sprintf("email:%s", user.Email)
	if err := p.cache.Set(emailCacheKey, user, time.Minute*30); err != nil {
		hlog.Infof("Cache user %d failed: %v\n", user.ID, err)
	}

	return nil
}

func (p *userRepository) Update(user *model.User) error {
	if err := p.db.Save(user).Error; err != nil {
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return db.ErrDB
	}

	idCacheKey := fmt.Sprintf("id:%d", user.ID)
	if err := p.cache.Delete(idCacheKey); err != nil {
		hlog.Infof("Delete user cache %d failed: %v\n", user.ID, err)
	}

	nameCacheKey := fmt.Sprintf("name:%s", user.Name)
	if err := p.cache.Delete(nameCacheKey); err != nil {
		hlog.Infof("Delete user cache %s failed: %v\n", user.Name, err)
	}

	emailCacheKey := fmt.Sprintf("email:%s", user.Email)
	if err := p.cache.Delete(emailCacheKey); err != nil {
		hlog.Infof("Delete user cache %s failed: %v\n", user.Email, err)
	}

	return nil
}
