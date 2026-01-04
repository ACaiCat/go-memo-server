package handler

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/pkg/mw"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type userRefreshReq struct {
	// Authorization 认证字符串
	Authorization string `header:"Authorization,required" vd:"regexp('^Bearer\\s+([A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*)$')" binding:"required" json:"Authorization"`
}

type userRefreshResp struct {
	// Token JWT令牌
	Token string `json:"token"`
}

// RefreshToken 刷新令牌
// @Summary      刷新用户登录令牌
// @Description  刷新用户的JWT令牌
// @Tags         认证
// @Accept       json
// @Produce      json
// @Param        request header userRefreshReq true "认证请求头"
// @Success      200  {object}  handler.BaseResp[userRefreshResp]  "令牌刷新成功"
// @Failure      400  "请求参数错误或Token格式不正确"
// @Failure      500  "服务器内部错误"
// @Router       /api/auth/refresh [post]
func (h *AuthHandler) RefreshToken(ctx context.Context, c *app.RequestContext) {
	req := new(userRefreshReq)
	if err := c.BindAndValidate(req); err != nil {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[userRefreshResp]{
			Status: consts.StatusBadRequest,
			Msg:    err.Error(),
		})
		return
	}
	token, found := strings.CutPrefix(req.Authorization, "Bearer ")

	if !found || len(token) == 0 {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[userRefreshResp]{
			Status: consts.StatusBadRequest,
			Msg:    "miss token",
		})
		return
	}

	token, err := mw.RefreshJWT(token)
	if err != nil {
		if errors.Is(err, mw.ErrTokenInvalid) {
			c.AbortWithStatusJSON(consts.StatusForbidden, handler.BaseResp[userRefreshResp]{
				Status: consts.StatusForbidden,
				Msg:    err.Error(),
			})
			return
		}

		if errors.Is(err, mw.ErrTokenExpiredTooLong) {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, handler.BaseResp[userRefreshResp]{
				Status: consts.StatusUnauthorized,
				Msg:    err.Error(),
			})
			return
		}

		log.Printf("failed to refresh auth token: %v\n", err)
		c.AbortWithStatusJSON(consts.StatusInternalServerError, handler.BaseResp[userRefreshResp]{
			Status: consts.StatusInternalServerError,
			Msg:    "internal server error",
		})
		return
	}
	c.JSON(consts.StatusOK, handler.BaseResp[userRefreshResp]{
		Status: consts.StatusOK,
		Msg:    "success",
		Data: &userRefreshResp{
			Token: token,
		},
	})
}
