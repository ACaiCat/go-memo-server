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
	// Token JWT令牌
	Token string `json:"token"`
	// UserID 用户ID
	UserID uint `json:"user_id"`
}

// UserRegister 用户注册
// @Summary      用户注册
// @Description  处理用户注册并返回JWT令牌
// @Tags         认证
// @Accept       x-www-form-urlencoded
// @Produce      json
// @Param        request formData userRegisterReq true "注册请求参数"
// @Success      200       {object}  handler.BaseResp[userRegisterResp]  "注册成功"
// @Failure      400       "请求参数错误"
// @Failure      409       "邮箱或用户名已被使用"
// @Failure      500       "服务器内部错误"
// @Router       /api/auth/register [post]
func (h *AuthHandler) UserRegister(ctx context.Context, c *app.RequestContext) {
	req := new(userRegisterReq)
	if err := c.BindAndValidate(req); err != nil {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[userRegisterResp]{
			Status: consts.StatusBadRequest,
			Msg:    err.Error(),
		})
		return
	}

	token, userID, err := h.userService.Create(req.Email, req.Name, req.Password)
	if err != nil {
		if errors.Is(err, service.ErrEmailUsed) {
			c.AbortWithStatusJSON(consts.StatusConflict, handler.BaseResp[userRegisterResp]{
				Status: consts.StatusConflict,
				Msg:    service.ErrEmailUsed.Error(),
			})
			return
		}
		if errors.Is(err, service.ErrNameUsed) {
			c.AbortWithStatusJSON(consts.StatusConflict, handler.BaseResp[userRegisterResp]{
				Status: consts.StatusConflict,
				Msg:    service.ErrNameUsed.Error(),
			})
			return
		}
		if errors.Is(err, service.ErrInvalidEmail) {
			c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[userRegisterResp]{
				Status: consts.StatusBadRequest,
				Msg:    service.ErrInvalidEmail.Error(),
			})
			return

		}

		hlog.Errorf("Failed to create user: %v\n", err)
		c.AbortWithStatusJSON(consts.StatusInternalServerError, handler.BaseResp[userRegisterResp]{
			Status: consts.StatusInternalServerError,
			Msg:    "internal server error",
		})
		return

	}

	c.JSON(consts.StatusOK, handler.BaseResp[userRegisterResp]{
		Status: consts.StatusOK,
		Msg:    "success",
		Data: &userRegisterResp{
			Token:  token,
			UserID: userID,
		},
	})
}
