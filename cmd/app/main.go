package main

import (
	"fmt"
	"log"

	"github.com/Zyprush18/badmintonzz/internal/interface/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// use env
	if err:= godotenv.Load();err != nil {
		log.Fatal("Failed Load Env")
	}


	router := gin.Default()

	v1 := router.Group("/v1")
	http.RegisteRoute(v1)
	
	fmt.Println("🚀 server running on Port: 8080")

	router.Run()
}
