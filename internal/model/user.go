package model

type User struct {
	// ID 用户ID
	ID uint `gorm:"primarykey,index" json:"id"`
	// Name 用户名
	Name string `gorm:"index" json:"name"`
	// Email 邮箱
	Email string `gorm:"index" json:"email"`
	// Password 密码加盐哈希
	Password string `json:"password" `
}
