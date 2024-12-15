package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func setupLog() {
	file, err := os.OpenFile("go后端错误日志.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("日志生成错误: %v", err)
	}
	log.SetOutput(file)
}

func main() {
	setupLog()
	r := gin.Default()

	r.GET("/abcd", func(c *gin.Context) {
		log.Println("Ping endpoint hit")
		c.JSON(200, gin.H{"message": "1111"})
	})

	r.Run(":1222")
}
