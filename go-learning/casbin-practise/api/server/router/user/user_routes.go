package user

import "github.com/gin-gonic/gin"

func InitRoutes(ginEngine *gin.Engine) {
	userRoute := ginEngine.Group("/user")
	{
		//userRoute.GET("/", Health)
		userRoute.POST("/", CreateUser)
		userRoute.DELETE("/:uid", DeleteUser)
		userRoute.GET("/:uid", GetUser)
		userRoute.GET("/", ListUser)
		userRoute.PUT("/", UpdateUser)
	}

}
