package main

import (
	"log"
	"time"
	"web-of-gin/initialization"
	"web-of-gin/modle/login"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// web 应用全局操作对象
var m initialization.Manager

// web 应用的初始化操作
func init() {
	// 设置日志格式
	log.SetFlags(log.Lshortfile)

	m.Init()
}

func main() {
	// test db
	db := m.DB()
	if db == nil {
		log.Fatalln("db is nil")
	}

	user := login.User{
		UserName: "test", Passwd: "123456", CreateAt: time.Now(),
	}

	db.Create(&user) // 插入一条数据

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"db": user,
		})
	})
	r.Run(":8080")
}
