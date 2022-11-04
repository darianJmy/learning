package user

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/types"
	"github.com/darianJmy/learning/go-learning/casbin-practise/app/cmd"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(200, "success")
}

func CreateUser(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	if err := cmd.CoreV1.User().CreateUser(context.TODO(), &user); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	c.String(200, "success")
}

func DeleteUser(c *gin.Context) {
	var uri types.UserUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	if err := cmd.CoreV1.User().DeleteUser(context.TODO(), uri.UID); err != nil {
		c.JSON(400, gin.H{"status": "erro"})
		return
	}

	c.String(200, "success")
}

func GetUser(c *gin.Context) {
	var uri types.UserUri

	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	user, err := cmd.CoreV1.User().GetUser(context.TODO(), uri.UID)
	if err != nil {
		c.JSON(400, gin.H{"status": "erro"})
		return
	}
	c.JSON(200, user)
}
