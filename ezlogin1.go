package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var users = map[string]string{
	"admin": "wenrp123", // 用户名: 密码
	"user1": "123456",
}

func main() {
	r := gin.Default()

	// 使用 CORS 中间件
	r.Use(cors.Default()) // 默认允许所有域名

	// 登录接口
	r.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// 绑定 JSON 数据
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// 验证用户名和密码
		if pwd, exists := users[loginData.Username]; exists && pwd == loginData.Password {
			c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		}
	})

	// 启动服务器
	r.Run(":1222") // 监听 8080 端口
}
