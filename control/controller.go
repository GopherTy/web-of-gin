package control

import (
	"web-of-gin/initialization"
)

// 数据库对象
var (
	db = initialization.DB()
)

// Controller 应用功能的控制器
type Controller struct {
}
