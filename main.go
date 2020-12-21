package main

import (
	"os"

	_ "web-of-gin/init"
	"web-of-gin/module/configs"
	"web-of-gin/module/logger"
	"web-of-gin/router"
	"web-of-gin/utils"

	"github.com/gin-gonic/gin"
)

// 项目入口
func main() {
	// 配置对象
	cnf := configs.Instance()

	// 是否启用 Gin 日志输出
	if cnf.Server.LogsPath != "" {
		err := utils.CreatePath(cnf.Server.LogsPath)
		if err != nil {
			logger.Instance().Fatal(err.Error())
		}

		f, err := os.Create(cnf.Server.LogsPath)
		if err != nil {
			logger.Instance().Fatal(err.Error())
		}
		gin.DefaultWriter = f
	}

	// 配置项目是否为稳定版,SetMode函数应该在gin.Default之前调用。
	if cnf.Server.Release {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	// 路由功能注册
	var r router.Router
	r.Route(engine)

	// 验证服务器是否以HTTPS的方式启动
	if cnf.Server.CertFile != "" && cnf.Server.KeyFile != "" {
		engine.RunTLS(cnf.Server.Address, cnf.Server.CertFile, cnf.Server.KeyFile)
	} else {
		engine.Run(cnf.Server.Address)
	}
}
