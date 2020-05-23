package middleware

import (
	"net/http"
	"web-of-gin/db"
	"web-of-gin/model/users"

	"github.com/gin-gonic/gin"
)

// UserManage 用户管理
func UserManage() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := db.Engine()

		exists, err := db.IsTableExist(&users.User{})
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
		}
		if !exists {
			db.CreateTables(&users.User{})
		}

		exists, err = db.IsTableExist(&users.UserInfo{})
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
		}
		if !exists {
			db.CreateTables(&users.UserInfo{})
		}

		exists, err = db.IsTableExist(&users.UserLoginLog{})
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
		}
		if !exists {
			db.CreateTables(&users.UserLoginLog{})
		}

		// before

		c.Next()

		//after 在中间件调用了处理函数之后，就会在此处调用
	}
}
