package router

import (
	"web-of-gin/control"

	"github.com/gin-gonic/gin"
)

// Router 路由管理
type Router struct {
}

// Init 注册路由
func (Router) Init(engine *gin.Engine) {
	// 控制器
	var ctl control.Controller

	// 用户入口
	engine.POST("/login", ctl.Login)
}
