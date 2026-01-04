package router

import (
	authHandler "github.com/ACaiCat/memo/internal/handler/auth"
	memoHandler "github.com/ACaiCat/memo/internal/handler/memo"
	"github.com/ACaiCat/memo/internal/repository"
	"github.com/ACaiCat/memo/internal/service"
	"github.com/ACaiCat/memo/pkg/mw"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func SetupRouter(r *server.Hertz, db *gorm.DB, cache *redis.Client) {
	userRepo := repository.NewUserRepository(db, cache)
	userService := service.NewUserService(userRepo)
	userHandler := authHandler.NewAuthHandler(userService)

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", userHandler.UserRegister)
		auth.POST("/login", userHandler.UserLogin)
		auth.POST("/refresh", userHandler.RefreshToken)
	}

	momoRepo := repository.NewMemoRepository(db, cache)
	momoService := service.NewMemoService(momoRepo)
	momoHandler := memoHandler.NewMemoHandler(momoService)

	memo := api.Group("/users/:user_id/memos")
	{
		memo.Use(mw.JWTAuth)
		memo.POST("/create", momoHandler.MemoCreate)
		memo.POST("/mark", momoHandler.MemoMark)
		memo.GET("/query", momoHandler.MemoQuery)
		memo.POST("/delete", momoHandler.MemoDelete)
	}
}
