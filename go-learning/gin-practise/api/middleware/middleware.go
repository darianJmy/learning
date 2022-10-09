package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/log"
)

func LoggerToFile() gin.HandlerFunc {
	handlerFunc := func(c *gin.Context) {
		startTime := time.Now()

		// 处理请求操作
		c.Next()

		endTime := time.Now()

		latencyTime := endTime.Sub(startTime)

		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()

		log.ConsoleInfo.Printf("| %3d | %13v | %15s | %s | %s |", statusCode, latencyTime, clientIp, reqMethod, reqUri)
	}
	return handlerFunc
}
