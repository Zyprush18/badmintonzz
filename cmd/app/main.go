package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	{
		v1 := router.Group("/v1")
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	fmt.Println("🚀 server running on Port: 8080")

	router.Run()
}
