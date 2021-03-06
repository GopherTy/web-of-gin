package test

import (
	"github.com/gin-gonic/gin"
)

// Dispatcher test功能包下的控制器，由全局控制器控制（Controller）控制。
type Dispatcher struct {
}

// Test 项目搭建的测试API
func (Dispatcher) Test(c *gin.Context) {
	c.JSON(200, "Hello World")
}

// MiddlewareHello 用户管理中间件的测试API
func (Dispatcher) MiddlewareHello(c *gin.Context) {
	c.JSON(200, "hello middleware  use success")
}
