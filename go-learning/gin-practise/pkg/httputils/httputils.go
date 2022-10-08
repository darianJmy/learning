package httputils

import "github.com/gin-gonic/gin"

func Success(c *gin.Context, msg interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  msg,
	})
}

func Failed(c *gin.Context, msg interface{}) {
	c.JSON(400, gin.H{
		"code": 400,
		"msg":  msg,
	})
}
