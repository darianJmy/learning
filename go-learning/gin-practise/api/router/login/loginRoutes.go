package login

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/learning/go-learning/gin-practise/api/middleware"
)

func InitRoutes(ginEngine *gin.Engine) {
	ginEngine.Use(middleware.LoggerToFile())
	userRoute := ginEngine.Group("/login")
	{
		userRoute.GET("/", Login)

	}
}
