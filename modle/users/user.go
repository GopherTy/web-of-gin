package users

import "time"

// User  系统注册用户登录表
type User struct {
	ID       int64  `gorm:"primary_key"`
	UserName string `gorm:"size:16"`
	Passwd   string
	CreateAt time.Time
	UpdateAt time.Time
}

// TableName 表名
func (User) TableName() string {
	return "user"
}
