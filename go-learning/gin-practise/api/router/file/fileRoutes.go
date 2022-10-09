package file

import "github.com/gin-gonic/gin"

func InitRoutes(ginEngine *gin.Engine) {
	userRoute := ginEngine.Group("/file")
	{
		userRoute.GET("/upload", UpLoad)
	}
}
