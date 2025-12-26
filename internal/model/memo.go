package model

import (
	"time"
)

type Memo struct {
	// ID 备忘录ID
	ID uint `gorm:"primarykey,index" json:"id"`
	// CreatedAt 创建时间
	CreatedAt time.Time ` json:"created_at"`
	// UserID 用户ID
	UserID uint `gorm:"index" json:"user_id"`
	// Title 标题
	Title string `gorm:"index" json:"title"`
	// Content 内容
	Content string `json:"content"`
	// Status 备忘录状态
	Status Status `json:"status"`
	// StartTime 开始时间
	StartTime time.Time `json:"start_time"`
	// EndTime 结束时间
	EndTime time.Time `json:"end_time"`
}
