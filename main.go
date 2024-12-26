package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files" // 引入 swaggerFiles 包
	"github.com/swaggo/gin-swagger"
	_ "mymodule/docs" // 引入Swagger文档生成的文件
)

//	@title			这是我学习API的swag文档
//	@version		6666
//	@description	This is a666
//	@termsOfService	666

//	@contact.name	非常666的文档
//	@contact.url	666
//	@contact.email	support@swagger.io

//	@host		http://127.0.0.1:1222
//	@BasePath	???

func main() {
	r := gin.Default()

	// Swagger 相关配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handlerfiles.Handler)) // 使用 files.Handler 代替 swaggerFiles.Handler

	// Example route
	r.GET("/hello", HelloHandler)

	r.Run(":8080")
}

// HelloHandler godoc
//
//	@Summary		jjj
//	@Description	sdss
//	@Tags			hi
//	@Accept			json
//	@Produce		json
//	@Success		200	{string}	string	"hello world"
//	@Router			/hello [get]
func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}
