package login

// UserInfo 用户信息表
type UserInfo struct {
	ID       int64 `gorm:"AUTO_INCREMENT"`
	UID      int64
	NickName string
	Age      int
	Phone    int
	Email    string
	Address  string
}

// TableName 表名
func (UserInfo) TableName() string {
	return "user_info"
}
