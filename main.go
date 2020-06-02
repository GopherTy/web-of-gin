package main

import (
	"flag"
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

var releaseMode bool

// 应用对象的初始化操作
func init() {
	flag.BoolVar(&releaseMode, "r", false, "set application mode to release,default value is false. example: -r or -r=true")
	flag.Parse()

	initialization.Init()
}

// 项目入口
func main() {
	// 配置对象
	cnf := config.Configure()

	// 是否启用 Gin 日志输出
	if cnf.Logger.GinLogsPath != "" {
		err := utils.CreatePath(cnf.Logger.GinLogsPath)
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}

		f, err := os.Create(cnf.Logger.GinLogsPath)
		if err != nil {
			logger.Logger().Fatal(err.Error())
		}
		gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
	}

	// 配置项目是否为稳定版,SetMode函数应该在gin.Default之前调用。
	if releaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()

	// 路由功能注册
	var r router.Router
	r.Route(engine)

	// 验证服务器是否以HTTPS的方式启动
	if cnf.HTTP.TLS {
		engine.RunTLS(cnf.HTTP.Address, cnf.HTTP.CertFile, cnf.HTTP.KeyFile)
	} else {
		engine.Run(cnf.HTTP.Address)
	}
}
