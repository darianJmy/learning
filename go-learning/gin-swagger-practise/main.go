package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strconv"

	_ "gin-swagger-practise/docs"
)

type user struct {
	Name string
	Age  int
}

//	@title			接囗文档
//	@version		1.0
//	@description	统一登陆项目
//	@termsofservice	https://github.com
//	@contact.name	JiXingXing
//	@contact.email	542255405@qq.com
//	@host			127.0.0.1:8080
//	@BasePath		/
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/user/:name", GetUser)
	r.GET("/user", ListUser)
	r.POST("/user", CreateUser)
	r.PUT("/user/:name", UpdateUser)
	r.DELETE("/user/:name", DeleteUser)

	r.Run(":8080")
}

//	@Summary		查看用户详细信息
//	@Description	查看用户详细信息
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"用户名称"
//	@Param			age		query		string	true	"用户年龄"
//	@Success		200		{object}	user	"请求成功"
//	@Router			/user/{name} [get]
func GetUser(c *gin.Context) {
	name := c.Param("name")
	a := c.Query("age")
	age, _ := strconv.Atoi(a)
	c.JSON(200, &user{Name: name, Age: age})
}

//	@Summary		查看用户列表
//	@Description	查看用户列表
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]user	"请求成功"
//	@Router			/user [get]
func ListUser(c *gin.Context) {
	c.JSON(200, &[]user{{Name: "jixingxing", Age: 18}, {Name: "yanglizi", Age: 17}})
}

//	@Summary		创建用户
//	@Description	创建用户
//	@Accept			json
//	@Produce		json
//	@Param			data	body		user	true	"请示参数data"
//	@Success		200		{object}	user	"请求成功"
//	@Router			/user [post]
func CreateUser(c *gin.Context) {
	var user user
	c.BindJSON(&user)
	c.JSON(200, &user)
}

//	@Summary		更新用户
//	@Description	更新用户
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"用户名称"
//	@Param			data	body		user	true	"请示参数data"
//	@Success		200		{object}	user	"请求成功"
//	@Router			/user/{name} [put]
func UpdateUser(c *gin.Context) {
	name := c.Param("name")
	var user user
	c.BindJSON(&user)
	user.Name = name
	c.JSON(200, &user)
}

//	@Summary		删除用户
//	@Description	删除用户
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"请示参数data"
//	@Success		200		{string}	string	"请求成功"
//	@Router			/user/{name} [delete]
func DeleteUser(c *gin.Context) {
	_ = c.Param("name")

	c.String(200, "Success")
}
