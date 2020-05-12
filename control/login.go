package control

import (
	"github.com/gin-gonic/gin"
)

// Login 登录到 web 系统
func (Controller) Login(c *gin.Context) {
	// db.Get(&users.User{})
	c.String(200, "success")
}
