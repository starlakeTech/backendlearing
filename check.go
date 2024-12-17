package main

import (
	"fmt"
	"net/http"
	"os"	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	if _, err := os.Stat("/www/wwwroot/aaaa/check.go"); os.IsNotExist(err) {
		fmt.Println("admin.html not found!")
	} else {
		fmt.Println("找到了")
	}

	if _, err := os.Stat("/www/wwwroot/html/user.html"); os.IsNotExist(err) {
		fmt.Println("user.html not found!")
	} else {
		fmt.Println("找到了")
	}

	
	r := gin.Default()

	r.LoadHTMLGlob("./www/wwwroot/html/*")
	r.Use(cors.Default())

	
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})

	
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", nil)
	})

	r.Run(":1222")
}
