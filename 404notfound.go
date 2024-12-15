package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Page notfound"})
	})

	r.Run(":1222")
}
