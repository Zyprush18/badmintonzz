package main

import (
	"fmt"
	"log"

	"github.com/Zyprush18/badmintonzz/internal/infrastucture/db"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// use env
	if err:= godotenv.Load();err != nil {
		log.Fatal("Failed Load Env")
	}

	if err:= db.Connect_DB();err != nil {
		log.Println(err.Error())
		log.Fatal("Failed to connect to database")
	}


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
