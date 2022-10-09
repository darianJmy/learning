package user

import "github.com/gin-gonic/gin"

func InitRoutes(ginEngine *gin.Engine) {
	userRoute := ginEngine.Group("/user")
	{
		userRoute.GET("/:username", GetUser)
		userRoute.GET("/list", ListUser)
		userRoute.DELETE("/:username", DeleteUser)
		userRoute.POST("/", CreateUser)
		userRoute.PUT("/:username", UpdateUser)
	}
}
