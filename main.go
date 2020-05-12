package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"web-of-gin/config"
	"web-of-gin/initialization"
	"web-of-gin/router"
	"web-of-gin/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// 应用的初始化操作
func init() {
	initialization.Init()
}

// 项目入口
func main() {
	fmt.Println("step 0")

	cfg := config.Configure()
	path := cfg.BasePath() + "/"

	var err error
	// 设置 Gin 日志输出
	if !utils.IsFileOrDirExists(path + "log") {
		err = os.Mkdir(path+"log", os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}

	f, err := os.Create(path + "log/gin.log")
	if err != nil {
		log.Fatalln(err)
	}
	w := io.MultiWriter(f, os.Stdout)

	gin.DefaultWriter = w

	engine := gin.Default()

	// 路由功能注册
	var r router.Router
	r.Init(engine)

	// 验证服务器是否以HTTPS的方式启动
	if cfg.HTTP.TLS {
		engine.RunTLS(cfg.HTTP.Address, path+cfg.HTTP.CertFile, path+cfg.HTTP.KeyFile)
	} else {
		engine.Run(cfg.HTTP.Address)
	}

}
