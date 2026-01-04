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

type MemoRepository interface {
	GetByID(id uint) (*model.Memo, error)
	GetByUserID(uid uint) ([]model.Memo, error)
	Search(uid uint, keywords []string, status model.Status, page int, perPage int) ([]model.Memo, int, error)
	Create(memo *model.Memo) error
	Update(memos []*model.Memo) error
	Delete(memos []*model.Memo) error
}

type memoRepository struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewMemoRepository(db *gorm.DB, redis *redis.Client) MemoRepository {
	cacheClient := cache.NewCache(redis, "memo")
	return &memoRepository{
		db:    db,
		cache: cacheClient,
	}
}

func (m *memoRepository) GetByID(id uint) (*model.Memo, error) {
	var memo model.Memo
	cacheKey := fmt.Sprintf("id:%d", id)

	if err := m.cache.Get(cacheKey, &memo); err == nil {
		hlog.Debugf("Hit memo cache: %d\n", id)
		return &memo, nil
	}

	if err := m.db.Where("id = ?", id).First(&memo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := m.cache.SetEmpty(cacheKey, time.Minute*5); err != nil {
				hlog.Errorf("Cache empty memo %d failed: %v\n", id, err)
			}
			return nil, nil
		}
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return nil, db.ErrDB
	}

	if err := m.cache.Set(cacheKey, &memo, time.Minute*30); err != nil {
		hlog.Errorf("Cache memo %d failed: %v\n", id, err)
	}

	return &memo, nil

}

func (m *memoRepository) GetByUserID(uid uint) ([]model.Memo, error) {
	var memos []model.Memo
	cacheKey := fmt.Sprintf("uid:%d", uid)

	if err := m.cache.Get(cacheKey, &memos); err == nil {
		hlog.Debugf("Hit user memos cache: %d\n", uid)
		return memos, nil
	}

	if err := m.db.Where("user_id = ?", uid).Find(&memos).Error; err != nil {
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return memos, db.ErrDB
	}

	if len(memos) == 0 {
		if err := m.cache.SetEmpty(cacheKey, time.Minute*5); err != nil {
			hlog.Errorf("Cache empty user memos %d failed: %v\n", uid, err)
		}
		return memos, nil
	}

	if err := m.cache.Set(cacheKey, &memos, time.Minute*30); err != nil {
		hlog.Errorf("Cache user memos %d failed: %v\n", uid, err)
	}

	return memos, nil
}

func (m *memoRepository) Search(uid uint, keywords []string, status model.Status, page int, perPage int) ([]model.Memo, int, error) {
	var memos []model.Memo
	tx := m.db.Model(&model.Memo{}).Where("user_id = ?", uid)

	if status != model.StatusAny {
		tx = tx.Where("status = ?", status)
	}

	for _, keyword := range keywords {
		tx = tx.Where("title LIKE ?", "%"+keyword+"%")
	}

	var total int64

	if err := tx.Count(&total).Error; err != nil {
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return memos, 0, db.ErrDB
	}

	if err := tx.Offset(perPage * (page - 1)).Limit(perPage).Find(&memos).Error; err != nil {
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return memos, 0, db.ErrDB
	}

	return memos, int(total), nil
}

func (m *memoRepository) Create(memo *model.Memo) error {
	if err := m.db.Create(memo).Error; err != nil {
		hlog.Errorf("%s: %v", db.ErrDB, err)
		return db.ErrDB
	}

	cacheKey := fmt.Sprintf("id:%d", memo.ID)
	if err := m.cache.Set(cacheKey, memo, time.Minute*30); err != nil {
		hlog.Infof("Cache memo %d failed: %v\n", memo.ID, err)
	}
	m.InvalidateUserMemoCache(memo.UserID)

	return nil

}

func (m *memoRepository) Update(memos []*model.Memo) error {
	tx := m.db.Begin()

	for _, memo := range memos {
		if err := tx.Save(memo).Error; err != nil {
			tx.Rollback()
			hlog.Errorf("%s: %v", db.ErrDB, err)
			return db.ErrDB
		}
	}

	if err := tx.Commit().Error; err != nil {
		hlog.Errorf("Commit transaction failed: %v", err)
		return db.ErrDB
	}

	cacheKeys := make([]string, 0, len(memos))
	userIDs := make(map[uint]bool)

	for _, memo := range memos {
		cacheKey := fmt.Sprintf("id:%d", memo.ID)
		cacheKeys = append(cacheKeys, cacheKey)
		userIDs[memo.UserID] = true
	}

	if err := m.cache.Delete(cacheKeys...); err != nil {
		hlog.Errorf("Delete memos cache failed: %v", err)
	}

	for userID := range userIDs {
		m.InvalidateUserMemoCache(userID)
	}

	return nil
}

func (m *memoRepository) Delete(memos []*model.Memo) error {
	tx := m.db.Begin()

	for _, memo := range memos {
		if err := tx.Delete(memo).Error; err != nil {
			tx.Rollback()
			hlog.Errorf("%s: %v", db.ErrDB, err)
			return db.ErrDB
		}
	}

	if err := tx.Commit().Error; err != nil {
		hlog.Errorf("Commit transaction failed: %v", err)
		return db.ErrDB
	}

	cacheKeys := make([]string, 0, len(memos))
	userIDs := make(map[uint]bool)

	for _, memo := range memos {
		cacheKey := fmt.Sprintf("id:%d", memo.ID)
		cacheKeys = append(cacheKeys, cacheKey)
		userIDs[memo.UserID] = true
	}

	if err := m.cache.Delete(cacheKeys...); err != nil {
		hlog.Errorf("Delete memos cache failed: %v", err)
	}

	for userID := range userIDs {
		m.InvalidateUserMemoCache(userID)
	}

	return nil
}

func (m *memoRepository) InvalidateUserMemoCache(uid uint) {
	cacheKey := fmt.Sprintf("uid:%d", uid)
	_ = m.cache.Delete(cacheKey)
}
