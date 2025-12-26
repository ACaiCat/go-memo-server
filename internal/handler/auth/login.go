package handler

import (
	"context"
	"errors"
	"log"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/internal/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type userLoginReq struct {
	// Email 用户邮箱
	Email string `form:"email,required" vd:"email($)" json:"email"`
	// Password 用户密码
	Password string `form:"password,required" vd:"len($)>4 && len($)<36" json:"password"`
}

type userLoginResp struct {
	handler.BaseResp
	// Token JWT令牌
	Token string `json:"token"`
}

// UserLogin 用户登录
// @Summary      用户登录
// @Description  处理用户登录并返回JWT令牌
// @Tags         认证
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        request formData userLoginReq true "登录请求参数"
// @Success      200       {object}  userLoginResp  "登录成功"
// @Failure      400       "请求参数错误"
// @Failure      403       "邮箱或密码错误"
// @Failure      500       "服务器内部错误"
// @Router       /api/auth/login [post]
func (h *AuthHandler) UserLogin(ctx context.Context, c *app.RequestContext) {
	resp := new(userLoginResp)
	req := new(userLoginReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}

	token, err := h.userService.ValidaUser(req.Email, req.Password)
	resp.Token = token
	if err != nil {
		if errors.Is(err, service.ErrPasswordError) || errors.Is(err, service.ErrUserNotFound) {
			log.Printf("failed to validate user: %v\n", err)
			resp.Status = consts.StatusForbidden
			resp.Msg = "email or password error"
			c.JSON(resp.Status, resp)
			return
		}
		resp.Status = consts.StatusInternalServerError
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}
	resp.Status = consts.StatusOK
	resp.Msg = "success"
	c.JSON(resp.Status, resp)
}
