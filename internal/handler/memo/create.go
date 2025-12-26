package handler

import (
	"context"
	"time"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/internal/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type memoCreateReq struct {
	// UserID 用户ID
	UserID uint `path:"user_id,required" swaggerignore:"true"`
	// Title 标题
	Title string `header:"title,required" vd:"len($)>0 && len($)< 100"  binding:"required" json:"title"`
	// Content 内容
	Content string `header:"content,required" vd:"len($)>0 && len($)< 5000"  binding:"required" json:"content"`
	// StartTimestamp 开始时间戳
	StartTimestamp int64 `header:"start_timestamp,required" vd:"$>=1577836800000 && $<=1893456000000"  binding:"required" json:"start_timestamp"`
	// EndTimestamp 结束时间戳
	EndTimestamp int64 `header:"end_timestamp,required" vd:"$>=1577836800000 && $<=1893456000000"  binding:"required" json:"end_timestamp"`
}

type memoCreateReqResp struct {
	handler.BaseResp
	// Memo 备忘录
	Memo *model.Memo `json:"memo"`
}

// MemoCreate 创建备忘录
// @Summary      创建备忘录
// @Description  创建一条新的备忘录
// @Tags         备忘录
// @Accept       json
// @Produce      json
// @Param        request header  memoCreateReq true "创建请求参数"
// @Param        user_id   path      int  true  "用户ID"
// @Success      200       {object}  memoCreateReqResp  "创建成功"
// @Failure      500       "服务器内部错误"
// @Security     ApiKeyAuth
// @Router       /api/users/{user_id}/memos/create [post]
func (h *MemoHandler) MemoCreate(ctx context.Context, c *app.RequestContext) {
	resp := new(memoCreateReqResp)
	req := new(memoCreateReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}

	startTime := time.UnixMilli(req.StartTimestamp)
	endTime := time.UnixMilli(req.EndTimestamp)

	memo, err := h.memoService.Create(req.UserID, req.Title, req.Content, startTime, endTime)
	if err != nil {
		hlog.Errorf("failed to create memo: %v\n", err)
		resp.Status = consts.StatusInternalServerError
		resp.Msg = "internal server error"
		c.JSON(resp.Status, resp)
		return

	}

	resp.Memo = memo
	resp.Msg = "success"
	resp.Status = consts.StatusOK
	c.JSON(resp.Status, resp)

}
