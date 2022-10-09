package user

import (
	"github.com/gin-gonic/gin"

	"github.com/darianJmy/learning/go-learning/gin-practise/cmd"
	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/httputils"
	"github.com/darianJmy/learning/go-learning/gin-practise/pkg/types"
)

func GetUser(c *gin.Context) {
	var uri types.UserUri
	if err := c.ShouldBindUri(&uri); err != nil {
		httputils.Failed(c, "Please enter the correct format")
		return
	}
	user, err := cmd.CoreV1.Mysql().GetUser(uri.UserName)
	if err != nil {
		httputils.Failed(c, "Failed to query the database")
		return
	}

	httputils.Success(c, user)
}

func ListUser(c *gin.Context) {
	users, err := cmd.CoreV1.Mysql().ListUser()
	if err != nil {
		httputils.Failed(c, "Failed to query the database")
		return
	}

	httputils.Success(c, users)
}

func DeleteUser(c *gin.Context) {
	var uri types.UserUri

	if err := c.ShouldBindUri(&uri); err != nil {
		httputils.Failed(c, "Please enter the correct format")
		return
	}

	if err := cmd.CoreV1.Mysql().DeleteUser(uri.UserName); err != nil {
		httputils.Failed(c, "Failed to delete to database")
		return
	}

	httputils.Success(c, "Success to delete data in database")
}

func CreateUser(c *gin.Context) {
	var user types.User
	if err := c.ShouldBindJSON(&user); err != nil {
		httputils.Failed(c, "Please enter the correct format")
		return
	}

	if err := cmd.CoreV1.Mysql().CreateUser(&user); err != nil {
		httputils.Failed(c, "Failed to write to database")
		return
	}

	httputils.Success(c, "Success to Insert data in database")
}

func UpdateUser(c *gin.Context) {
	var uri types.UserUri
	var user types.User

	if err := c.ShouldBindUri(&uri); err != nil {
		httputils.Failed(c, "Please enter the correct format")
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		httputils.Failed(c, "Please enter the correct format")
		return
	}

	if err := cmd.CoreV1.Mysql().UpdateUser(uri.UserName, &user); err != nil {
		httputils.Failed(c, "Failed to write to database")
		return
	}

	httputils.Success(c, "Success to Update data in database")
}
