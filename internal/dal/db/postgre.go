package db

import (
	"errors"

	"github.com/ACaiCat/memo/pkg/config"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
import "gorm.io/driver/postgres"

var (
	ErrDB = errors.New("db: db operation failed")
)

func InitPostgre() *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.GetConfig().DSN),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	)
	if err != nil {
		hlog.Fatalf("Failed to connect DB: %v", err)
	}
	return db
}
