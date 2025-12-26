package handler

import "github.com/ACaiCat/memo/internal/service"

type MemoHandler struct {
	memoService service.MemoService
}

func NewMemoHandler(memoService service.MemoService) *MemoHandler {
	return &MemoHandler{
		memoService: memoService,
	}
}
