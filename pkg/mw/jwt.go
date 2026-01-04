package mw

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/ACaiCat/memo/internal/handler"
	"github.com/ACaiCat/memo/pkg/config"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenExpired        = errors.New("token is expired")
	ErrTokenInvalid        = errors.New("token is invalid")
	ErrTokenExpiredTooLong = errors.New("token is expired too long")
	ErrUserNotMatch        = errors.New("token does not belong to the specified user")
)

type jwtAuthReq struct {
	UserID        uint   `path:"user_id,required"`
	Authorization string `header:"Authorization,required" vd:"regexp('^Bearer\\s+([A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*)$')"`
}

func JWTAuth(ctx context.Context, c *app.RequestContext) {
	var req jwtAuthReq
	if err := c.BindAndValidate(&req); err != nil {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[any]{
			Status: consts.StatusBadRequest,
			Msg:    err.Error(),
		})
		return
	}

	authHeader := req.Authorization
	token, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found || len(token) == 0 {
		c.AbortWithStatusJSON(consts.StatusBadRequest, handler.BaseResp[any]{
			Status: consts.StatusBadRequest,
			Msg:    "miss token",
		})
		return
	}
	err := ValidateJWT(req.UserID, token)
	if err != nil {
		if errors.Is(err, ErrTokenExpired) {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, handler.BaseResp[any]{
				Status: consts.StatusUnauthorized,
				Msg:    err.Error(),
			})
			return
		}
		if errors.Is(err, ErrUserNotMatch) {
			c.AbortWithStatusJSON(consts.StatusUnauthorized, handler.BaseResp[any]{
				Status: consts.StatusUnauthorized,
				Msg:    err.Error(),
			})
			return
		}
		if errors.Is(err, ErrTokenInvalid) {
			c.AbortWithStatusJSON(consts.StatusForbidden, handler.BaseResp[any]{
				Status: consts.StatusForbidden,
				Msg:    err.Error(),
			})
			return
		}
		log.Printf("failed to validate user jwt: %v", err)
		c.AbortWithStatusJSON(consts.StatusInternalServerError, handler.BaseResp[any]{
			Status: consts.StatusInternalServerError,
			Msg:    err.Error(),
		})
		return
	}
	c.Set("user_id", req.UserID)
	c.Next(ctx)
}

func NewJWT(uid uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": strconv.Itoa(int(uid)),
		"exp":     time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	signedJWT, err := token.SignedString([]byte(config.GetConfig().ServerConfig.JWTSecrete))
	if err != nil {
		return "", err
	}

	return signedJWT, nil

}

func ValidateJWT(uid uint, jwrStr string) error {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwrStr, &claims, func(token *jwt.Token) (any, error) {
		return []byte(config.GetConfig().ServerConfig.JWTSecrete), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return ErrTokenExpired
		}
		return ErrTokenInvalid
	}

	if userID, ok := claims["user_id"]; ok {
		if userID != strconv.Itoa(int(uid)) {
			return ErrUserNotMatch
		}

		return nil
	}

	return ErrTokenInvalid

}

func RefreshJWT(jwrStr string) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwrStr, &claims, func(token *jwt.Token) (any, error) {
		return []byte(config.GetConfig().ServerConfig.JWTSecrete), nil
	}, jwt.WithExpirationRequired())
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return "", ErrTokenInvalid
	}

	expVal, _ := claims["exp"]

	exp, _ := expVal.(float64)

	if time.Now().After(time.Unix(int64(exp), 0).Add(7 * 24 * time.Hour)) {
		return "", ErrTokenExpiredTooLong
	}

	if userIDVal, ok := claims["user_id"]; ok {
		userID, err := strconv.Atoi(userIDVal.(string))
		if err != nil {
			return "", err
		}
		token, err := NewJWT(uint(userID))
		if err != nil {
			return "", err
		}

		return token, nil

	}

	return "", ErrTokenInvalid
}
