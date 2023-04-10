## Gin-Swagger 使用说明

### 安装

1. 本地需要有swag命令
2. 下载swagger代码
3. 需要通过swag生成docs文件夹下文件

### 下载swag命令与项目
```
# 安装swag命令，如果安装完成没有，就去 $gopath/bin 下看看
go install github.com/swaggo/swag/cmd/swag@latest
    
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles
```

### 引用swagger
```
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

### 编写代码生成器文件
```
### 具体内容可以百度搜索
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

//	@Summary		查看用户详细信息
//	@Description	查看用户详细信息
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"用户名称"
//	@Param			age		query		string	true	"用户年龄"
//	@Success		200		{object}	user	"请求成功"
//	@Router			/user/{name} [get]
```

### 执行swag命令
```
### 格式化代码生成器内容
swag fmt

### 生成代码放到docs下
swag init
```

### 执行代码，打开swagger UI
```
http://127.0.0.1:8080/swagger/index.html
```

