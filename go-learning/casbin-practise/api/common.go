package api

import (
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/server/router/user"
	"github.com/gin-gonic/gin"
)

func RegistryRoutes(ginEngine *gin.Engine) {

	user.InitRoutes(ginEngine)
}
