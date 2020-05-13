package control

import (
	"github.com/gin-gonic/gin"
)

// Test 项目搭建的测试API
func (Controller) Test(c *gin.Context) {
	c.JSON(200, "Hello World")
}

// UserMiddlewareTest 用户管理中间件的测试API
func (Controller) UserMiddlewareTest(c *gin.Context) {
	result, _ := c.Get("result")

	if v, ok := result.(string); ok {
		c.JSON(200, v)
	} else {
		c.JSON(200, "key is not exists")
	}
}
