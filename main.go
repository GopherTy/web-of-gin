package main

import (
	"io"
	"os"
	"web-of-gin/config"
	"web-of-gin/initialization"
	"web-of-gin/logger"
	"web-of-gin/router"
	"web-of-gin/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 应用对象的初始化操作
func init() {
	initialization.Init()
}

// 项目入口
func main() {
	// 配置对象
	cfg := config.Configure()
	path := utils.BasePath() + "/"

	// 是否启用 Gin 日志输出
	if cfg.Logger.OutputLogs {
		f, err := os.Create(path + "log/gin.log")
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
		gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	}

	engine := gin.Default()

	// 路由功能注册
	var r router.Router
	r.Route(engine)

	// 验证服务器是否以HTTPS的方式启动
	if cfg.HTTP.TLS {
		engine.RunTLS(cfg.HTTP.Address, path+cfg.HTTP.CertFile, path+cfg.HTTP.KeyFile)
	} else {
		engine.Run(cfg.HTTP.Address)
	}
}
