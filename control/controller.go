package control

import (
	"web-of-gin/initialization"

	"github.com/go-xorm/xorm"
)

// Controller 应用功能的控制器
type Controller struct {
}

// DB 数据库对象
func (Controller) DB() (db *xorm.Engine) {
	db = initialization.DB()
	return
}
