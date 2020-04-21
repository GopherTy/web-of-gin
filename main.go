package main

import (
	"log"
	"web-of-gin/initialization"
	"web-of-gin/modle/login"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// web 应用的初始化操作
func init() {
	initialization.Init()
}

func main() {
	// test db
	db := initialization.DB()
	if db == nil {
		log.Fatalln("db is nil")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		user := login.User{}

		db.First(&user)
		user.Passwd = "123987"
		db.Save(&user)

		c.JSON(200, gin.H{
			"db": user,
		})
	})

	r.Run(":8080")
}
