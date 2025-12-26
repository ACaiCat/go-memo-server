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

type userRegisterReq struct {
	// Name 用户名
	Name string `form:"name,required" vd:"len($)>2 && len($)<10" binding:"required" json:"name"`
	// Email 用户邮箱
	Email string `form:"email,required" vd:"email($)" binding:"required" json:"email"`
	// Password 用户密码
	Password string `form:"password,required" vd:"len($)>4 && len($)<36" binding:"required" json:"password"`
}

type userRegisterResp struct {
	handler.BaseResp
	// Token JWT令牌
	Token string `json:"token"`
}

// UserRegister 用户注册
// @Summary      用户注册
// @Description  处理用户注册并返回JWT令牌
// @Tags         认证
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        request formData userRegisterReq true "注册请求参数"
// @Success      200       {object}  userRegisterResp  "注册成功"
// @Failure      400       "请求参数错误"
// @Failure      409       "邮箱或用户名已被使用"
// @Failure      500       "服务器内部错误"
// @Router       /api/auth/register [post]
func (h *AuthHandler) UserRegister(ctx context.Context, c *app.RequestContext) {
	resp := new(userRegisterResp)
	req := new(userRegisterReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}

	token, err := h.userService.Create(req.Email, req.Name, req.Password)
	resp.Token = token
	if err != nil {
		if errors.Is(err, service.ErrEmailUsed) {
			resp.Status = consts.StatusConflict
			resp.Msg = service.ErrEmailUsed.Error()
			c.JSON(resp.Status, resp)
			return
		}
		if errors.Is(err, service.ErrNameUsed) {
			resp.Status = consts.StatusConflict
			resp.Msg = service.ErrNameUsed.Error()
			c.JSON(resp.Status, resp)
			return
		}
		if errors.Is(err, service.ErrInvalidEmail) {
			resp.Status = consts.StatusBadRequest
			resp.Msg = service.ErrInvalidEmail.Error()
			c.JSON(resp.Status, resp)
			return

		}

		hlog.Errorf("Failed to create user: %v\n", err)
		resp.Status = consts.StatusInternalServerError
		resp.Msg = "internal server error"
		c.JSON(resp.Status, resp)
		return

	}

	resp.Msg = "success"
	resp.Status = consts.StatusOK
	c.JSON(resp.Status, resp)

}
