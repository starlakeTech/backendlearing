package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var users = map[string]string{
	"admin": "wenrp123", //
	"user1": "123456",
}

func main() {
	r := gin.Default()

	r.Use(cors.Default()) // 不加cors这个传不了数据

	// 登录接口
	r.POST("/login", func(c *gin.Context) {
		var loginData struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// 绑定 JSON 数据
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "请求失败"})
			return
		}

		// 验证
		if pwd, exists := users[loginData.Username]; exists && pwd == loginData.Password {
			c.JSON(http.StatusOK, gin.H{"message": "登陆成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "用户名或密码错误"})
		}
	})

	r.Run(":1222")
}
