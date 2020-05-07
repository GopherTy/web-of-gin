package users

// UserLoginLog  用户登录系统日志表
type UserLoginLog struct {
	ID       int64 `gorm:"AUTO_INCREMENT"`
	UID      int64
	ClientIP string
}

// TableName 表名
func (UserLoginLog) TableName() string {
	return "user_logs"
}
