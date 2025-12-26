package service

import (
	"errors"
	"strings"
	"time"

	"github.com/ACaiCat/memo/internal/model"
	"github.com/ACaiCat/memo/internal/repository"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var (
	ErrMemoNotFound     = errors.New("service: memo not found")
	ErrMemoCreate       = errors.New("service: memo create failed")
	ErrMemoGet          = errors.New("service: memo get failed")
	ErrMemoUpdate       = errors.New("service: memo update failed")
	ErrMemoDelete       = errors.New("service: memo delete failed")
	ErrMemoSearch       = errors.New("service: memo search failed")
	ErrMemoNotBelong    = errors.New("service: memo not belong to current user")
	ErrNotSupportStatus = errors.New("service: status is not supported")
)

type MemoService interface {
	Create(uid uint, title string, content string, startTime time.Time, endTime time.Time) (*model.Memo, error)
	Mark(uid uint, memoIDs []uint, status model.Status) error
	Search(uid uint, query string, status model.Status, page int, perPage int) (*[]model.Memo, int, error)
	Delete(uid uint, memoIDs []uint) error
}

func NewMemoService(repo repository.MemoRepository) MemoService {
	return &memoService{repo: repo}
}

type memoService struct {
	repo repository.MemoRepository
}

func (m *memoService) Create(uid uint, title string, content string, startTime time.Time, endTime time.Time) (*model.Memo, error) {

	memo := &model.Memo{
		UserID:    uid,
		Title:     title,
		Content:   content,
		Status:    model.StatusPending,
		CreatedAt: time.Now(),
		StartTime: startTime,
		EndTime:   endTime,
	}

	err := m.repo.Create(memo)
	if err != nil {
		hlog.Errorf("%s: %s", ErrMemoCreate, err)
		return nil, ErrMemoCreate
	}

	return memo, nil
}

func (m *memoService) Mark(uid uint, memoIDs []uint, status model.Status) error {
	var memos []*model.Memo

	if status != model.StatusCompleted && status != model.StatusPending {
		return ErrNotSupportStatus
	}

	for _, id := range memoIDs {
		memo, err := m.repo.GetByID(id)
		if err != nil {
			hlog.Errorf("%s: %s", ErrMemoGet, err)
			return ErrMemoGet
		}
		if memo == nil {
			return ErrMemoNotFound
		}
		if memo.UserID != uid {
			return ErrMemoNotBelong
		}
		memo.Status = status
		memos = append(memos, memo)
	}

	err := m.repo.Update(memos)
	if err != nil {
		hlog.Errorf("%s: %s", ErrMemoUpdate, err)
		return ErrMemoUpdate
	}

	return nil
}

func (m *memoService) Search(uid uint, query string, status model.Status, page int, perPage int) (*[]model.Memo, int, error) {
	if status != model.StatusCompleted && status != model.StatusPending && status != model.StatusAny {
		return nil, 0, ErrNotSupportStatus
	}

	keywords := strings.Split(query, " ")

	memos, total, err := m.repo.Search(uid, keywords, status, page, perPage)
	if err != nil {
		hlog.Errorf("%s: %s", ErrMemoSearch, err)
		return nil, 0, ErrMemoSearch
	}
	return &memos, total, nil
}

func (m *memoService) Delete(uid uint, memoIDs []uint) error {
	var memos []*model.Memo

	for _, id := range memoIDs {
		memo, err := m.repo.GetByID(id)
		if err != nil {
			hlog.Errorf("%s: %s", ErrMemoGet, err)
			return ErrMemoGet
		}
		if memo == nil {
			return ErrMemoNotFound
		}
		if memo.UserID != uid {
			return ErrMemoNotBelong
		}
		memos = append(memos, memo)
	}

	err := m.repo.Delete(memos)
	if err != nil {
		hlog.Errorf("%s: %s", ErrMemoDelete, err)
		return ErrMemoDelete
	}

	return nil
}
