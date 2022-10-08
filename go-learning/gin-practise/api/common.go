package api

import (
	"github.com/darianJmy/learning/go-learning/gin-practise/api/router/user"
	"github.com/gin-gonic/gin"
)

func RegistryRoutes(ginEngine *gin.Engine) {

	user.InitRoutes(ginEngine)
}
