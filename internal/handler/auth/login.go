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
	// Token JWT令牌
	Token string `json:"token"`
	// UserID 用户ID
	UserID uint `json:"user_id"`
}

// UserLogin 用户登录
// @Summary      用户登录
// @Description  处理用户登录并返回JWT令牌
// @Tags         认证
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        request formData userLoginReq true "登录请求参数"
// @Success      200       {object}  handler.BaseResp[userLoginResp]  "登录成功"
// @Failure      400       "请求参数错误"
// @Failure      403       "邮箱或密码错误"
// @Failure      500       "服务器内部错误"
// @Router       /api/auth/login [post]
func (h *AuthHandler) UserLogin(ctx context.Context, c *app.RequestContext) {
	req := new(userLoginReq)
	if err := c.BindAndValidate(req); err != nil {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[userLoginResp]{
			Status: consts.StatusBadRequest,
			Msg:    err.Error(),
		})
		return
	}

	token, userID, err := h.userService.ValidaUser(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrPasswordError) || errors.Is(err, service.ErrUserNotFound) {
			log.Printf("failed to validate user: %v\n", err)
			c.AbortWithStatusJSON(consts.StatusForbidden, handler.BaseResp[userLoginResp]{
				Status: consts.StatusForbidden,
				Msg:    "email or password error",
			})
			return
		}
		c.AbortWithStatusJSON(consts.StatusInternalServerError, handler.BaseResp[userLoginResp]{
			Status: consts.StatusInternalServerError,
			Msg:    err.Error(),
		})
		return
	}
	c.JSON(consts.StatusInternalServerError, handler.BaseResp[userLoginResp]{
		Status: consts.StatusOK,
		Msg:    "success",
		Data: &userLoginResp{
			Token:  token,
			UserID: userID,
		},
	})
}
