package login

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/httputils"
)

func Login(c *gin.Context) {

	username := c.Query("username")
	password := c.Query("password")
	if username == "" || password == "" {
		httputils.Failed(c, "Failed to query the database")
		return
	}
	httputils.Success(c, gin.H{"username": username, "password": password})
}
