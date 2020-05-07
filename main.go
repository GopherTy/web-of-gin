package main

import (
	"web-of-gin/control"
	"web-of-gin/initialization"
	"web-of-gin/router"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// web 应用的初始化操作
func init() {
	initialization.Init()
}

// 项目入口
func main() {
	engine := gin.Default()

	var r router.Router
	r.Regist("/ping", router.MethodFunc{Method: "get", Function: control.Login}) // 注册路由

	r.Init(engine)
	engine.Run(":8080")
}
