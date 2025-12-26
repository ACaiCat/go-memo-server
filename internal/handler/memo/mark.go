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
	handler.BaseResp
}

// MemoMark 标记备忘录
// @Summary      标记备忘录
// @Description  标记备忘录的状态
// @Tags         备忘录
// @Accept       json
// @Produce      json
// @Param        request body  memoMarkReq true "标记请求参数"
// @Param        user_id   path      int  true  "用户ID"
// @Success      200       {object}  memoMarkResp  "修改成功"
// @Failure      400       "备忘录状态无效"
// @Failure      500       "服务器内部错误"
// @Security     ApiKeyAuth
// @Router       /api/users/{user_id}/memos/mark [post]
func (h *MemoHandler) MemoMark(ctx context.Context, c *app.RequestContext) {
	resp := new(memoMarkResp)
	req := new(memoMarkReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}

	err := h.memoService.Mark(req.UserID, req.MemoIDs, req.Status)
	if err != nil {
		if errors.Is(err, service.ErrNotSupportStatus) {
			resp.Status = consts.StatusBadRequest
			resp.Msg = service.ErrNotSupportStatus.Error()
			c.JSON(resp.Status, resp)
			return
		}

		if errors.Is(err, service.ErrMemoNotFound) {
			resp.Status = consts.StatusBadRequest
			resp.Msg = service.ErrMemoNotFound.Error()
			c.JSON(resp.Status, resp)
			return
		}

		hlog.Errorf("Failed to mark memo: %v\n", err)
		resp.Status = consts.StatusInternalServerError
		resp.Msg = "internal server error"
		c.JSON(resp.Status, resp)
		return
	}

	resp.Msg = "success"
	resp.Status = consts.StatusOK
	c.JSON(resp.Status, resp)

}
