package role

import (
	"context"
	"github.com/darianJmy/learning/go-learning/casbin-practise/api/types"
	"github.com/darianJmy/learning/go-learning/casbin-practise/app/cmd"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(200, "success")
}

func CreateRole(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	if err := cmd.CoreV1.User().CreateRole(context.TODO(), &user); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	c.String(200, "success")
}

func DeleteRole(c *gin.Context) {
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

func GetRole(c *gin.Context) {
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

func ListRole(c *gin.Context) {
	user, err := cmd.CoreV1.User().ListUser(context.TODO())
	if err != nil {
		c.JSON(400, gin.H{"status": "erro"})
		return
	}
	c.JSON(200, user)
}

func UpdateRole(c *gin.Context) {
	var user types.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"status": "error"})
		return
	}

	err := cmd.CoreV1.User().UpdateUser(context.TODO(), &user)
	if err != nil {
		c.JSON(400, gin.H{"status": "erro"})
		return
	}

	c.JSON(200, nil)

}
