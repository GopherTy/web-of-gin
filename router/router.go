package router

import (
	"web-of-gin/control"
	"web-of-gin/middleware"

	"github.com/gin-gonic/gin"
)

// Router 路由管理
type Router struct {
}

// Route 注册路由
func (Router) Route(engine *gin.Engine) {
	// 控制器
	var ctl control.Controller

	// 非用户组
	engine.GET("/", ctl.TestDispacher.Test)                                                         // 测试接口注册
	engine.GET("/middleware/test", middleware.HelloMiddleware(), ctl.TestDispacher.MiddlewareHello) // test 中间件

	// 用户组
	group := engine.Group("/app")
	group.GET("/test", ctl.TestDispacher.Test)
}
