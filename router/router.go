package router

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Router 路由管理
type Router struct {
	pathFunc map[string]MethodFunc // 用于存放路由关系的 map
}

// MethodFunc 方法函数
type MethodFunc struct {
	Method   string             //路由请求方法。
	Function func(*gin.Context) // 路由对应的功能函数
}

// Regist 路由注册管理
func (r *Router) Regist(path string, f MethodFunc) {
	if r.pathFunc == nil {
		r.pathFunc = make(map[string]MethodFunc)
	}

	r.pathFunc[path] = f
}

// Init 注册路由
func (r *Router) Init(engine *gin.Engine) {
	if r.pathFunc != nil {
		for path, v := range r.pathFunc {
			v.Method = strings.ToUpper(v.Method)
			switch v.Method {
			case "GET":
				engine.GET(path, v.Function)
			}
		}
	}
}
