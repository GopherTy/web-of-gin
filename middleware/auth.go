package middleware

import (
	"net/http"
	"web-of-gin/control"
	"web-of-gin/model/auth"
	"web-of-gin/model/users"

	"github.com/gin-gonic/gin"
)

// UserAuthenticate 用户登录认证
func UserAuthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctl control.Controller
		db := ctl.DB()

		exists, err := db.IsTableExist(&auth.Role{})
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
		}
		if !exists {
			db.CreateTables(&users.UserLoginLog{})
		}

		exists, err = db.IsTableExist(&auth.Auth{})
		if err != nil {
			c.AbortWithError(http.StatusBadGateway, err)
		}
		if !exists {
			db.CreateTables(&users.UserLoginLog{})
		}

		c.Next()
	}
}
