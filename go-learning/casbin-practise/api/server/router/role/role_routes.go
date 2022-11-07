package role

import "github.com/gin-gonic/gin"

func InitRoutes(ginEngine *gin.Engine) {
	userRoute := ginEngine.Group("/role")
	{
		//userRoute.GET("/", Health)
		userRoute.POST("/", CreateRole)
		userRoute.DELETE("/:uid", DeleteRole)
		userRoute.GET("/:uid", GetRole)
		userRoute.GET("/", ListRole)
		userRoute.PUT("/", UpdateRole)
	}

}
