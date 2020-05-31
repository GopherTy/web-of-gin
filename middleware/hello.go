package middleware

import (
	"web-of-gin/logger"

	"github.com/gin-gonic/gin"
)

// HelloMiddleware middleware defined
func HelloMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		// before use middleware
		logger.Logger().Info("before ...")

		c.Next()

		// after use middleware
		logger.Logger().Info("after ...")
	}
}
