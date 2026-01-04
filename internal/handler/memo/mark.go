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

type memoMarkReq struct {
	// UserID 用户ID
	UserID uint `path:"user_id,required" swaggerignore:"true"`
	// MemoIDs 备忘录ID
	MemoIDs []uint `json:"memo_ids,required" binding:"required"`
	// Status 状态
	Status model.Status `json:"status,required" binding:"required"`
}

type memoMarkResp struct {
}

// MemoMark 标记备忘录
// @Summary      标记备忘录
// @Description  标记备忘录的状态
// @Tags         备忘录
// @Accept       json
// @Produce      json
// @Param        request body  memoMarkReq true "标记请求参数"
// @Param        user_id   path      int  true  "用户ID"
// @Success      200       {object}  handler.BaseResp[memoMarkResp]  "修改成功"
// @Failure      400       "备忘录状态无效"
// @Failure      500       "服务器内部错误"
// @Security     ApiKeyAuth
// @Router       /api/users/{user_id}/memos/mark [post]
func (h *MemoHandler) MemoMark(ctx context.Context, c *app.RequestContext) {
	req := new(memoMarkReq)
	if err := c.BindAndValidate(req); err != nil {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[memoMarkResp]{
			Status: consts.StatusBadRequest,
			Msg:    err.Error(),
		})
		return
	}

	err := h.memoService.Mark(req.UserID, req.MemoIDs, req.Status)
	if err != nil {
		if errors.Is(err, service.ErrNotSupportStatus) {
			c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[memoMarkResp]{
				Status: consts.StatusBadRequest,
				Msg:    service.ErrNotSupportStatus.Error(),
			})
			return
		}

		if errors.Is(err, service.ErrMemoNotFound) {
			c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[memoMarkResp]{
				Status: consts.StatusBadRequest,
				Msg:    service.ErrMemoNotFound.Error(),
			})
			return
		}

		hlog.Errorf("Failed to mark memo: %v\n", err)
		c.AbortWithStatusJSON(consts.StatusInternalServerError, handler.BaseResp[memoMarkResp]{
			Status: consts.StatusInternalServerError,
			Msg:    "internal server error",
		})
		return
	}

	c.AbortWithStatusJSON(consts.StatusOK, handler.BaseResp[memoMarkResp]{
		Status: consts.StatusOK,
		Msg:    "success",
	})

}
