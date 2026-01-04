package handler

import (
	"context"
	"errors"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/internal/model"
	"github.com/ACaiCat/memo/internal/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type memoQueryReq struct {
	// UserID 用户ID
	UserID uint `path:"user_id,required" swaggerignore:"true"`
	// Query 关键词
	Query string `query:"query,required" default:"" json:"query" example:""`
	// Status 状态
	Status model.Status `query:"status" default:"0" binding:"required" json:"status" example:"0"`
	// Page 页码
	Page int `query:"page,required" binding:"required" json:"page" example:"1"`
	// PerPage 每页条目数
	PerPage int `query:"per_page,required" json:"per_page" example:"5"`
}

type memoQueryResp struct {
	// Memos 满足条件的备忘录
	Memos *[]model.Memo `json:"memos"`
}

// MemoQuery 查询备忘录
// @Summary      查询备忘录
// @Description  查询满足要求的备忘录
// @Tags         备忘录
// @Accept       json
// @Produce      json
// @Param        request query memoQueryReq true "查询请求参数"
// @Param        user_id   path      int  true  "用户ID"
// @Success      200       {object}  handler.BaseResp[memoMarkResp]  "查询成功"
// @Failure      400       "备忘录状态无效"
// @Failure      500       "服务器内部错误"
// @Security     ApiKeyAuth
// @Router       /api/users/{user_id}/memos/query [get]
func (h *MemoHandler) MemoQuery(ctx context.Context, c *app.RequestContext) {
	req := new(memoQueryReq)
	if err := c.BindAndValidate(req); err != nil {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[memoQueryResp]{
			Status: consts.StatusBadRequest,
			Msg:    err.Error(),
		})
		return
	}

	memos, total, err := h.memoService.Search(req.UserID, req.Query, req.Status, req.Page, req.PerPage)
	if err != nil {
		if errors.Is(err, service.ErrNotSupportStatus) {
			c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[memoQueryResp]{
				Status: consts.StatusBadRequest,
				Msg:    service.ErrNotSupportStatus.Error(),
			})
			return
		}

		hlog.Errorf("Failed to search memo: %v\n", err)
		c.AbortWithStatusJSON(consts.StatusInternalServerError, handler.BaseResp[memoQueryResp]{
			Status: consts.StatusInternalServerError,
			Msg:    "internal server error",
		})
		return

	}

	c.JSON(consts.StatusOK, handler.BaseResp[memoQueryResp]{
		Status: consts.StatusOK,
		Msg:    "success",
		Data: &memoQueryResp{
			Memos: memos,
		},
		Pagination: &handler.Pagination{
			Page:    req.Page,
			PerPage: req.PerPage,
			Total:   total,
		},
	})
}
