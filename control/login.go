package control

import (
	"fmt"
	"net/http"
	"web-of-gin/initialization"
	"web-of-gin/model/users"

	"github.com/gin-gonic/gin"
)

// Login 登录到 web 系统
func Login(c *gin.Context) {
	db := initialization.DB()
	user := users.User{}
	fmt.Println("-------------")
	db.First(&user)

	c.JSON(http.StatusOK, user)

}
