package handler

import (
	"context"
	"errors"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/internal/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type memoDeleteReq struct {
	// UserID 用户ID
	UserID uint `path:"user_id,required" swaggerignore:"true"`
	// MemoIDs 备忘录ID
	MemoIDs []uint `json:"memo_ids,required" binding:"required"`
}

type memoDeleteResp struct {
	handler.BaseResp
}

// MemoDelete 删除备忘录
// @Summary      删除备忘录
// @Description  删除备忘录
// @Tags         备忘录
// @Accept       json
// @Produce      json
// @Param        request body  memoDeleteReq true "删除请求参数"
// @Param        user_id   path      int  true  "用户ID"
// @Success      200       {object}  memoDeleteResp  "删除成功"
// @Failure      500       "服务器内部错误"
// @Security     ApiKeyAuth
// @Router       /api/users/{user_id}/memos/delete [post]
func (h *MemoHandler) MemoDelete(ctx context.Context, c *app.RequestContext) {
	resp := new(memoDeleteResp)
	req := new(memoDeleteReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}

	err := h.memoService.Delete(req.UserID, req.MemoIDs)
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

		hlog.Errorf("Failed to delete memo: %v\n", err)
		resp.Status = consts.StatusInternalServerError
		resp.Msg = "internal server error"
		c.JSON(resp.Status, resp)
		return

	}

	resp.Msg = "success"
	resp.Status = consts.StatusOK
	c.JSON(resp.Status, resp)

}
