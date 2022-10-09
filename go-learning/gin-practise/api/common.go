package api

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/learning/go-learning/gin-practise/api/router/file"
	"github.com/darianJmy/learning/go-learning/gin-practise/api/router/login"
	"github.com/darianJmy/learning/go-learning/gin-practise/api/router/user"
)

func RegistryRoutes(ginEngine *gin.Engine) {

	user.InitRoutes(ginEngine)

	file.InitRoutes(ginEngine)

	login.InitRoutes(ginEngine)
}
