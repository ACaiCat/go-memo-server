package handler

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/internal/mw"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type userRefreshReq struct {
	// Authorization 认证字符串
	Authorization string `header:"Authorization,required" vd:"regexp('^Bearer\\s+([A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*)$')" binding:"required" json:"Authorization"`
}

type userRefreshResp struct {
	handler.BaseResp
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
// @Success      200  {object}  userRefreshResp  "令牌刷新成功"
// @Failure      400  "请求参数错误或Token格式不正确"
// @Failure      500  "服务器内部错误"
// @Router       /api/auth/refresh [post]
func (h *AuthHandler) RefreshToken(ctx context.Context, c *app.RequestContext) {
	resp := new(userRefreshResp)
	req := new(userRefreshReq)
	if err := c.BindAndValidate(req); err != nil {
		resp.Status = consts.StatusBadRequest
		resp.Msg = err.Error()
		c.JSON(resp.Status, resp)
		return
	}
	token, found := strings.CutPrefix(req.Authorization, "Bearer ")

	if !found || len(token) == 0 {
		resp.Status = consts.StatusBadRequest
		resp.Msg = "miss token"
		c.JSON(resp.Status, resp)
		return
	}

	token, err := mw.RefreshJWT(token)
	resp.Token = token
	if err != nil {
		if errors.Is(err, mw.ErrTokenInvalid) {
			resp.Status = consts.StatusForbidden
			resp.Msg = err.Error()
			c.JSON(resp.Status, resp)
			return
		}

		if errors.Is(err, mw.ErrTokenExpiredTooLong) {
			resp.Status = consts.StatusUnauthorized
			resp.Msg = err.Error()
			c.JSON(resp.Status, resp)
			return
		}

		log.Printf("failed to refresh auth token: %v\n", err)
		resp.Status = consts.StatusInternalServerError
		resp.Msg = "internal server error"
		c.JSON(resp.Status, resp)
		return
	}
	resp.Msg = "success"
	resp.Status = consts.StatusOK
	c.JSON(resp.Status, resp)
}
