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
	handler.BaseResp
	// Pagination 分页信息
	Pagination *handler.Pagination `json:"pagination"`
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
// @Success      200       {object}  memoQueryResp  "查询成功"
// @Failure      400       "备忘录状态无效"
// @Failure      500       "服务器内部错误"
// @Security     ApiKeyAuth
// @Router       /api/users/{user_id}/memos/query [get]
func (h *MemoHandler) MemoQuery(ctx context.Context, c *app.RequestContext) {
	resp := new(memoQueryResp)
	req := new(memoQueryReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}

	memos, total, err := h.memoService.Search(req.UserID, req.Query, req.Status, req.Page, req.PerPage)
	if err != nil {
		if errors.Is(err, service.ErrNotSupportStatus) {
			resp.Status = consts.StatusBadRequest
			resp.Msg = service.ErrNotSupportStatus.Error()
			c.JSON(resp.Status, resp)
			return
		}

		hlog.Errorf("failed to search memo: %v\n", err)
		resp.Status = consts.StatusInternalServerError
		resp.Msg = "internal server error"
		c.JSON(resp.Status, resp)
		return

	}

	resp.Memos = memos
	resp.Pagination = &handler.Pagination{
		Page:    req.Page,
		PerPage: req.PerPage,
		Total:   total,
	}
	resp.Msg = "success"
	resp.Status = consts.StatusOK
	c.JSON(resp.Status, resp)
}
